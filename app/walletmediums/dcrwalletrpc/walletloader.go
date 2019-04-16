package dcrwalletrpc

import (
	"context"
	"errors"
	"fmt"
	"github.com/decred/dcrd/dcrutil"
	"github.com/decred/dcrd/hdkeychain"
	"github.com/decred/dcrwallet/rpc/walletrpc"
	"github.com/decred/dcrwallet/walletseed"
	"github.com/raedahgroup/godcr/app"
	"github.com/raedahgroup/godcr/app/walletcore"
	"google.golang.org/grpc/codes"
)

func (c *WalletRPCClient) WalletExists() (bool, error) {
	res, err := c.walletLoader.WalletExists(context.Background(), &walletrpc.WalletExistsRequest{})
	if err != nil {
		return false, err
	}
	return res.Exists, nil
}

func (c *WalletRPCClient) GenerateNewWalletSeed() (string, error) {
	seed, err := hdkeychain.GenerateSeed(hdkeychain.RecommendedSeedLen)
	if err != nil {
		return "", err
	}

	return walletseed.EncodeMnemonic(seed), nil
}

func (c *WalletRPCClient) CreateWallet(passphrase, seed string) error {
	seedBytes, err := walletseed.DecodeUserInput(seed)
	if err != nil {
		return err
	}

	_, err = c.walletLoader.CreateWallet(context.Background(), &walletrpc.CreateWalletRequest{
		PrivatePassphrase: []byte(passphrase),
		Seed:              seedBytes,
	})

	// wallet will be opened if the create operation was successful
	if err == nil {
		c.walletOpen = true
	}

	return err
}

func (c *WalletRPCClient) OpenWalletIfExist(ctx context.Context) (walletExists bool, err error) {
	c.walletOpen = false
	loadWalletDone := make(chan bool)

	go func() {
		defer func() {
			loadWalletDone <- true
		}()

		walletExists, err = c.WalletExists()
		if err != nil || !walletExists {
			return
		}

		_, err = c.walletLoader.OpenWallet(context.Background(), &walletrpc.OpenWalletRequest{})

		// ignore wallet already open errors, it could be that dcrwallet loaded the wallet when it was launched by the user
		// or godcr opened the wallet without closing it
		if isRpcErrorCode(err, codes.AlreadyExists) {
			err = nil
		}
	}()

	select {
	case <-loadWalletDone:
		// if err is nil, then wallet was opened
		c.walletOpen = err == nil
		return

	case <-ctx.Done():
		return false, ctx.Err()
	}
}

func (c *WalletRPCClient) CloseWallet() {
	// don't actually close wallet loaded by dcrwallet
	// - if wallet wasn't opened by godcr, closing it could cause troubles for user
	// - even if wallet was opened by godcr, closing it without closing dcrwallet would cause troubles for user
	// when they next launch godcr
}

func (l *WalletRPCClient) DeleteWallet() error {
	return errors.New("wallet cannot be deleted when connecting via dcrwallet rpc")
}

func (c *WalletRPCClient) IsWalletOpen() bool {
	// for now, assume that the wallet's already open since we're connecting through dcrwallet daemon
	// ideally, we'd have to use dcrwallet's WalletLoaderService to do this
	return c.walletOpen
}

func (c *WalletRPCClient) SyncBlockChain(listener *app.BlockChainSyncListener, showLog bool) error {
	ctx := context.Background()

	bestBlockHeight, err := c.BestBlock()
	if err != nil {
		return err
	}

	syncStream, err := c.walletLoader.SpvSync(ctx, &walletrpc.SpvSyncRequest{})
	if err != nil {
		return err
	}

	// create wrapper around success listener and call rpc SubscribeToBlockNotifications
	// method associates the wallet with the consensus RPC server, subscribes the wallet for attached block and chain switch notifications,
	// and causes the wallet to process these notifications in the background.
	// also publish any pending transactions using PublishUnminedTransactions
	originalSyncEndedListener := listener.SyncEnded
	listener.SyncEnded = func(err error) {
		if err != nil {
			_, err := c.walletLoader.SubscribeToBlockNotifications(ctx, &walletrpc.SubscribeToBlockNotificationsRequest{})
			if err != nil {
				// no point pubslishing if above function did not succeed
				c.walletService.PublishUnminedTransactions(ctx, &walletrpc.PublishUnminedTransactionsRequest{})
			}
		}
		originalSyncEndedListener(err)
	}

	s := &spvSync{
		listener:  listener,
		netType:   c.NetType(),
		client:    syncStream,
		bestBlock: int64(bestBlockHeight),
	}

	// receive sync updates from stream and send to listener in separate goroutine
	go s.streamBlockchainSyncUpdates(showLog)
	return nil
}

func (c *WalletRPCClient) WalletConnectionInfo() (info walletcore.ConnectionInfo, err error) {
	accounts, loadAccountErr := c.AccountsOverview(walletcore.DefaultRequiredConfirmations)
	if loadAccountErr != nil {
		err = fmt.Errorf("error fetching account balance: %s", loadAccountErr.Error())
		info.TotalBalance = "0 DCR"
	} else {
		var totalBalance dcrutil.Amount
		for _, acc := range accounts {
			totalBalance += acc.Balance.Total
		}
		info.TotalBalance = totalBalance.String()
	}

	bestBlock, bestBlockErr := c.BestBlock()
	if bestBlockErr != nil && err != nil {
		err = fmt.Errorf("%s, error in fetching best block %s", err.Error(), bestBlockErr.Error())
	} else if bestBlockErr != nil {
		err = bestBlockErr
	}

	info.LatestBlock = bestBlock
	info.NetworkType = c.NetType()
	info.PeersConnected = c.GetConnectedPeersCount()

	return
}

func (c *WalletRPCClient) BestBlock() (uint32, error) {
	req, err := c.walletService.BestBlock(context.Background(), &walletrpc.BestBlockRequest{})
	if err != nil {
		return 0, err
	}
	return req.Height, err
}

func (c *WalletRPCClient) GetConnectedPeersCount() int32 {
	return numberOfPeers
}

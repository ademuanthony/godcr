package app

import (
	"context"
	"github.com/raedahgroup/godcr/app/sync"
	"github.com/raedahgroup/godcr/app/walletcore"
)

// WalletMiddleware defines key functions for interacting with a decred wallet
// These functions are implemented by the different mediums that provide access to a decred wallet
type WalletMiddleware interface {
	WalletExists() (bool, error)

	GenerateNewWalletSeed() (string, error)

	CreateWallet(passphrase, seed string) error

	SyncBlockChainOld(listener *BlockChainSyncListener, showLog bool) error

	SyncBlockChain(showLog bool, syncInfoUpdated func(privateSyncData *sync.PrivateInfo)) error

	RescanBlockChain() error

	// OpenWalletIfExist checks if the wallet the user is trying to access exists and opens the wallet
	// This method may stall if the wallet database is in use by some other process,
	// hence the need for ctx, so user can cancel the operation if it's taking too long
	// todo: some wallets may not use default public passphrase,
	// todo: in such cases request public passphrase from user to use in OpenWallet
	OpenWalletIfExist(ctx context.Context) (walletExists bool, err error)

	// CloseWallet is triggered whenever the godcr program is about to be terminated
	// Usually such termination attempts are halted to allow this function perform shutdown and cleanup operations
	CloseWallet()

	DeleteWallet() error

	IsWalletOpen() bool

	WalletConnectionInfo() (info walletcore.ConnectionInfo, err error)

	// BestBlock fetches the best block on the network
	BestBlock() (uint32, error)

	// GetConnectedPeersCount returns the number of connected peers
	GetConnectedPeersCount() int32

	walletcore.Wallet
}

// BlockChainSyncListener holds functions that are called during a blockchain sync operation to provide update on the sync operation
type BlockChainSyncListener struct {
	SyncStarted         func()
	SyncEnded           func(err error)
	OnHeadersFetched    func(percentageProgress int64)
	OnDiscoveredAddress func(state string)
	OnRescanningBlocks  func(percentageProgress int64)
	OnPeersUpdated      func(peerCount int32)
}

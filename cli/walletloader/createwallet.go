package walletloader

import (
	"context"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"

	"github.com/decred/dcrd/hdkeychain"
	"github.com/decred/dcrwallet/walletseed"
	"github.com/raedahgroup/godcr/app/config"
	"github.com/raedahgroup/godcr/app/walletmediums/dcrlibwallet"
	"github.com/raedahgroup/godcr/cli/termio"
	"github.com/raedahgroup/godcr/cli/termio/terminalprompt"
)

// createWallet creates a new wallet using the dcrlibwallet WalletMiddleware.
// User is prompted to select the network type for the wallet to be created.
// If no wallet for that type already exist, user is asked to provide a private passphrase for the wallet.
// After which, a new wallet seed is generated and shown to the user.
func createWallet(ctx context.Context, cfg *config.Config) (dcrlibwalletMiddleware *dcrlibwallet.DcrWalletLib, err error) {
	newWalletNetwork, err := requestNetworkTypeForNewWallet()
	if err != nil {
		return
	}

	// create dcrlibwallet wallet middleware and check if wallet of this type already exist
	dcrlibwalletMiddleware, err = prepareMiddlewareToCreateNewWallet(ctx, newWalletNetwork)
	if err != nil {
		return
	}

	newWalletPassphrase, err := requestNewWalletPassphrase()
	if err != nil {
		return
	}

	// get and display new wallet seed
	seed, err := generateNewWalletSeedAndDisplay()
	if err != nil {
		return
	}

	// user says they have backed up the generated wallet seed, finalize wallet creation
	err = dcrlibwalletMiddleware.CreateWallet(newWalletPassphrase, seed)
	if err != nil {
		return nil, fmt.Errorf("\nError creating wallet: %s.", err.Error())
	}
	fmt.Printf("Decred %s wallet created successfully at\n", dcrlibwalletMiddleware.NetType())
	fmt.Println(dcrlibwalletMiddleware.WalletDbDir)

	sync, err := runInitialSync(cfg)
	if err != nil || !sync {
		return dcrlibwalletMiddleware, err
	}

	return dcrlibwalletMiddleware, SyncBlockChain(ctx, dcrlibwalletMiddleware)
}

func requestNetworkTypeForNewWallet() (string, error) {
	// this function will be called when a user responds to the prompt to specify network type for new wallet
	checkNetworkTypeSelection := func(input string) error {
		if input == "" || // use default
			strings.EqualFold(input, "mainnet") || strings.EqualFold(input, "m") ||
			strings.EqualFold(input, "testnet") || strings.EqualFold(input, "t") ||
			strings.EqualFold(input, "simnet") || strings.EqualFold(input, "s") {
			return nil
		}
		return fmt.Errorf("invalid choice, please enter 'M' or 't' or 's'")
	}

	// prompt user to select network type for new wallet
	prompt := "Which net? (M)ainnet, (t)estnet, or (s)imnet? [M]"
	userResponse, err := terminalprompt.RequestInput(prompt, checkNetworkTypeSelection)
	if err != nil {
		return "", fmt.Errorf("\nError getting network type for new wallet: %s.", err.Error())
	}

	if strings.EqualFold(userResponse, "testnet") || strings.EqualFold(userResponse, "t") {
		return "testnet3", nil
	} else if strings.EqualFold(userResponse, "simnet") || strings.EqualFold(userResponse, "s") {
		return "simnet", nil
	}

	return "mainnet", nil
}

// prepareMiddlewareToCreateNewWallet reads appdata dir from godcr.conf to use as wallet db dir.
// Also ensures that a wallet of the specified type does not already exist in the appdata dir.
// If it exists, a new directory is created to hold the new wallet.
func prepareMiddlewareToCreateNewWallet(ctx context.Context, newWalletNetwork string) (*dcrlibwallet.DcrWalletLib, error) {
	// get appdata dir from config to place new wallet into
	cfg, err := config.ReadConfigFile()
	if err != nil {
		return nil, fmt.Errorf("\nError reading config file to determine directory to place new wallet: %s.", err.Error())
	}

	// find a suitable, unused dir to place new wallet
	var walletDbDir string
	networkDir := newWalletNetwork
	networkDirSuffix := 0
	for {
		walletDbDir = filepath.Join(cfg.AppDataDir, networkDir)
		_, err := os.Stat(walletDbDir)
		if err != nil && os.IsNotExist(err) {
			break
		}

		networkDirSuffix++
		networkDir = fmt.Sprintf("%s-%d", newWalletNetwork, networkDirSuffix)
	}

	return dcrlibwallet.Connect(ctx, walletDbDir, newWalletNetwork)
}

// requestNewWalletPassphrase asks user to enter private passphrase for new wallet twice.
// Prompt is repeated if both entered passphrases don't match.
func requestNewWalletPassphrase() (string, error) {
	for {
		passphrase, err := terminalprompt.RequestInputSecure("Enter private passphrase for new wallet", terminalprompt.EmptyValidator)
		if err != nil {
			return "", fmt.Errorf("\nError reading new wallet passphase: %s.", err.Error())
		}
		confirmPassphrase, err := terminalprompt.RequestInputSecure("Confirm passphrase", terminalprompt.EmptyValidator)
		if err != nil {
			return "", fmt.Errorf("\nError reading new wallet confirm passphase: %s.", err.Error())
		}

		if passphrase != confirmPassphrase {
			fmt.Println("Passphrases don't match, try again.")
			continue
		}

		return passphrase, nil
	}
}

func generateNewWalletSeedAndDisplay() (string, error) {
	// generate seed
	seed, err := hdkeychain.GenerateSeed(hdkeychain.RecommendedSeedLen)
	if err != nil {
		return "", fmt.Errorf("\nError generating seed for new wallet: %s.", err)
	}
	seedWords := walletseed.EncodeMnemonic(seed)

	// display seed
	fmt.Println("Your wallet generation seed is:")
	fmt.Println("-------------------------------")

	allWords := strings.Split(seedWords, " ")
	maxWordCountPerColumn := int(math.Ceil(float64(len(allWords)) / 3.0))
	col1Words := allWords[:maxWordCountPerColumn]
	col2Words := allWords[maxWordCountPerColumn : maxWordCountPerColumn*2]
	col3Words := allWords[maxWordCountPerColumn*2:]

	stdOutUsingTabs := termio.TabWriter(os.Stdout)
	for i := range col1Words {
		word1 := fmt.Sprintf("(%d)%s", i+1, col1Words[i])
		word2 := fmt.Sprintf("(%d)%s", i+maxWordCountPerColumn+1, col2Words[i])
		var word3 string
		if i < len(col3Words) {
			word3 = fmt.Sprintf("(%d)%s", i+(maxWordCountPerColumn*2)+1, col3Words[i])
		}
		fmt.Fprintf(stdOutUsingTabs, "%s\t%s\t%s\n", word1, word2, word3)
	}
	stdOutUsingTabs.Flush()

	fmt.Printf("Hex: %x\n", seed)
	fmt.Println("-------------------------------")

	return seedWords, nil
}

func runInitialSync(cfg *config.Config) (bool, error) {
	if cfg.InterfaceMode != "cli" {
		// do not attempt to sync on cli if the user requested a different interface to be launched
		// all other interfaces perform sync on launch
		return false, nil
	}

	if cfg.SyncBlockchain {
		// no need to ask user if to sync since `--sync` was already specified
		return true, nil
	} else {
		syncBlockchainPrompt := "Would you like to sync the blockchain now?"
		syncBlockchain, err := terminalprompt.RequestYesNoConfirmation(syncBlockchainPrompt, "Y")
		if err != nil {
			return false, fmt.Errorf("\nError reading your response: %s.", err.Error())
		}
		return syncBlockchain, nil
	}
}

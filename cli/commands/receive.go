package commands

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/mdp/qrterminal"
	"github.com/raedahgroup/godcr/app/walletcore"
	"github.com/raedahgroup/godcr/cli/runner"
	"github.com/raedahgroup/godcr/cli/termio/terminalprompt"
)

// ReceiveCommand generates an address for a user to receive DCR.
type ReceiveCommand struct {
	runner.WalletCommand
	Args ReceiveCommandArgs `positional-args:"yes"`
}
type ReceiveCommandArgs struct {
	AccountName string `positional-arg-name:"account-name"`
}

// Run runs the `receive` command.
func (receiveCommand ReceiveCommand) Run(ctx context.Context, wallet walletcore.Wallet, args []string) error {
	var accountNumber uint32
	// if no account name was passed in
	if receiveCommand.Args.AccountName == "" {
		// display menu options to select account
		var err error
		accountNumber, err = selectAccount(wallet)
		if err != nil {
			return err
		}
	} else {
		// if an account name was passed in e.g. ./godcr receive default
		// get the address corresponding to the account name and use it
		var err error
		accountNumber, err = wallet.AccountNumber(receiveCommand.Args.AccountName)
		if err != nil {
			return fmt.Errorf("Error fetching account number: %s", err.Error())
		}
	}

	receiveAddress, err := wallet.GenerateReceiveAddress(accountNumber)
	if err != nil {
		return err
	}

	// Print out address as string
	fmt.Println(receiveAddress)

	// Print out QR code
	validateConfirm := func(userResponse string) error {
		userResponse = strings.TrimSpace(userResponse)
		userResponse = strings.Trim(userResponse, `"`)
		if userResponse == "" || strings.EqualFold("Y", userResponse) || strings.EqualFold("n", userResponse) {
			return nil
		} else {
			return fmt.Errorf("invalid option, try again")
		}
	}

	confirm, err := terminalprompt.RequestInput("Would you like to generate a QR code? (y/N)", validateConfirm)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading your response: %s", err.Error())
		return err
	}

	if strings.EqualFold(confirm, "y") {
		qrterminal.GenerateHalfBlock("https://github.com/mdp/qrterminal", qrterminal.L, os.Stdout)

	}

	return nil
}
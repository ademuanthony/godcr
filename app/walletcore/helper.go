package walletcore

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/decred/dcrd/chaincfg"
	"github.com/decred/dcrd/dcrutil"
	"github.com/decred/dcrd/txscript"
	"github.com/raedahgroup/dcrlibwallet/txhelper"
)

type SyncStatus uint8

const (
	SyncStatusNotStarted SyncStatus = iota
	SyncStatusSuccess
	SyncStatusError
	SyncStatusInProgress
)

// GetAddressFromPkScript extracts the address from the supplied pkScript in the given chaincfg params
func GetAddressFromPkScript(activeNet *chaincfg.Params, pkScript []byte) (address string, err error) {
	_, addresses, _, err := txscript.ExtractPkScriptAddrs(txscript.DefaultScriptVersion,
		pkScript, activeNet)
	if err != nil {
		return
	}
	if len(addresses) < 1 {
		return "", errors.New("Cannot extract any address from output")
	}

	encodedAddresses := make([]string, len(addresses))
	for i, address := range addresses {
		encodedAddresses[i] = address.EncodeAddress()
	}

	return strings.Join(encodedAddresses, ", "), nil
}

func SimpleBalance(balance *Balance, detailed bool) string {
	if detailed {
		return balance.Total.String()
	} else {
		return balance.String()
	}
}

func SpendableBalance(balance *Balance) string {
	return balance.Spendable.String()
}

// NormalizeBalance adds 0s to the left and the right of balanceString to make it xxxx.xxxxxxxx DCR
func NormalizeBalance(balanceString string) string {
	parts := strings.Split(balanceString, " ")
	if !strings.Contains(parts[0], ".") {
		parts[0] = parts[0] + ".0"
	}
	numberParts := strings.Split(parts[0], ".")
	if len(numberParts[1]) < 8 {
		numberParts[1] = padRight(numberParts[1], "0", 8)
	}
	return fmt.Sprintf("%s.%s %s", numberParts[0], numberParts[1], parts[1])
}

func padRight(input, niddle string, maxLen int) string {
	for len(input) < maxLen {
		input += niddle
	}
	return input
}

// GetChangeDestinationsWithRandomAmounts generates change destination(s) based on the number of change address the user want
func GetChangeDestinationsWithRandomAmounts(wallet Wallet, nChangeOutputs int, amountInAtom int64, sourceAccount uint32,
	nUtxoSelection int, sendDestinations []txhelper.TransactionDestination) (changeOutputDestinations []txhelper.TransactionDestination, err error) {

	var changeAddresses []string
	for i := 0; i < nChangeOutputs; i++ {
		address, err := wallet.GenerateNewAddress(sourceAccount)
		if err != nil {
			return nil, fmt.Errorf("error generating address: %s", err.Error())
		}
		changeAddresses = append(changeAddresses, address)
	}

	changeAmount, err := txhelper.EstimateChange(nUtxoSelection, amountInAtom, sendDestinations, changeAddresses)
	if err != nil {
		return nil, fmt.Errorf("error in getting change amount: %s", err.Error())
	}
	if changeAmount <= 0 {
		return
	}

	var portionRations []float64
	var rationSum float64
	for i := 0; i < nChangeOutputs; i++ {
		portion := rand.Float64()
		portionRations = append(portionRations, portion)
		rationSum += portion
	}

	for i, portion := range portionRations {
		portionPercentage := portion / rationSum
		amount := portionPercentage * float64(changeAmount)

		changeOutput := txhelper.TransactionDestination{
			Address: changeAddresses[i],
			Amount:  dcrutil.Amount(amount).ToCoin(),
		}
		changeOutputDestinations = append(changeOutputDestinations, changeOutput)
	}
	return
}

func BuildTxDestinations(destinationAddresses []string, destinationAmounts []string) (destinations []txhelper.TransactionDestination, err error) {
	destinations = make([]txhelper.TransactionDestination, len(destinationAddresses))
	for i := range destinationAddresses {
		amount, err := strconv.ParseFloat(destinationAmounts[i], 64)
		if err != nil {
			return destinations, err
		}
		destinations[i] = txhelper.TransactionDestination{
			Address: destinationAddresses[i],
			Amount:  amount,
		}
	}
	return
}

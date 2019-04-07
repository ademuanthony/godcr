package walletcore

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/decred/dcrd/dcrutil"
	"github.com/raedahgroup/dcrlibwallet/txhelper"
)

type SyncStatus uint8

const (
	SyncStatusNotStarted SyncStatus = iota
	SyncStatusSuccess
	SyncStatusError
	SyncStatusInProgress
)

const StoreSeedWarningText = "Keep the seed in a safe place as you will NOT be able to restore your wallet without it. " +
	"Please keep in mind that anyone who has access to the seed can also restore your wallet " +
	"thereby giving them access to all your funds, so it is imperative that you keep it in a secure location."

func SimpleBalance(balance *Balance, detailed bool) string {
	if detailed {
		return balance.Total.String()
	} else {
		return balance.String()
	}
}

// NormalizeBalance adds 0s the right of balance to make it x.xxxxxxxx DCR
func NormalizeBalance(balance float64) string {
	return fmt.Sprintf("%010.8f DCR", balance)
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

func GetAllUtxos(walletMiddleware Wallet, accountNumber uint32, requiredConfirmations int32) (utxos []string, total int64, err error) {
	allUtxos, err := walletMiddleware.UnspentOutputs(accountNumber, 0, requiredConfirmations)
	if err != nil {
		return nil, 0, err
	}

	var totalInputAmountDcr float64
	for _, utxo := range allUtxos {
		utxos = append(utxos, utxo.OutputKey)
		totalInputAmountDcr += utxo.Amount.ToCoin()
	}

	totalInputAmount, err := dcrutil.NewAmount(totalInputAmountDcr)
	if err != nil {
		return
	}

	total = int64(totalInputAmount)
	return
}

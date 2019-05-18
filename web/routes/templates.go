package routes

import (
	"fmt"
	"github.com/raedahgroup/dcrlibwallet/txhelper"
	"html/template"
	"strings"
	"time"

	"github.com/decred/dcrd/dcrutil"
	"github.com/raedahgroup/godcr/app/utils"
	"github.com/raedahgroup/godcr/app/walletcore"
)

type templateData struct {
	name string
	path string
}

func templates() []templateData {
	return []templateData{
		{"error.html", "web/views/error.html"},
		{"createwallet.html", "web/views/createwallet.html"},
		{"overview.html", "web/views/overview.html"},
		{"sync.html", "web/views/sync.html"},
		{"send.html", "web/views/send.html"},
		{"receive.html", "web/views/receive.html"},
		{"history.html", "web/views/history.html"},
		{"transaction_details.html", "web/views/transaction_details.html"},
		{"staking.html", "web/views/staking.html"},
		{"accounts.html", "web/views/accounts.html"},
		{"security.html", "web/views/security.html"},
		{"settings.html", "web/views/settings.html"},
	}
}

func templateFuncMap() template.FuncMap {
	return template.FuncMap{
		"simpleBalance": func(balance *walletcore.Balance, detailed bool) string {
			if detailed {
				return walletcore.NormalizeBalance(balance.Total.ToCoin())
			}
			return balance.String()
		},
		"spendableBalance": func(balance *walletcore.Balance) string {
			return walletcore.NormalizeBalance(balance.Spendable.ToCoin())
		},
		"splitBalanceIntoParts": func(accounts []*walletcore.Account) []string {
			var totalBalance float64
			for _, account := range accounts {
				totalBalance += account.Balance.Total.ToCoin()
			}

			return utils.SplitAmountIntoParts(totalBalance)
		},
		"splitAmountIntoParts": func(amount int64) []string {
			dcrAmount := dcrutil.Amount(amount).ToCoin()

			return utils.SplitAmountIntoParts(dcrAmount)
		},
		"intSum": func(numbers ...int) (sum int) {
			for _, n := range numbers {
				sum += n
			}
			return
		},
		"accountString": func(account *walletcore.Account) string {
			if account.Balance.Unconfirmed > 0 {
				return fmt.Sprintf("%s %s (unconfirmed %s)", account.Name,
					walletcore.NormalizeBalance(account.Balance.Total.ToCoin()), walletcore.NormalizeBalance(account.Balance.Unconfirmed.ToCoin()))
			}
			return fmt.Sprintf("%s %s", account.Name, walletcore.NormalizeBalance(account.Balance.Total.ToCoin()))
		},
		"noUnconfirmedBalance": func(accounts []*walletcore.Account) bool {
			for _, account := range accounts {
				if account.Balance.Unconfirmed > 0 {
					return false
				}
			}
			return true
		},
		"amountDcr": func(amount int64) string {
			return dcrutil.Amount(amount).String()
		},
		"timestamp": func() int64 {
			return time.Now().Unix()
		},
		"extractDateTime": func(timestamp int64) string {
			utcTime := time.Unix(timestamp, 0).UTC()
			return utcTime.Format("2006-01-02 15:04:05")
		},
		"truncate": func(input string, maxLength int) string {
			if len(input) <= maxLength {
				return input
			}
			return fmt.Sprintf("%s...", input[0:maxLength])
		},
		"accountName" : func(txn *walletcore.Transaction) string {
			var accountNames []string
			isInArray := func(accountName string) bool {
				for _, name := range accountNames {
					if name == accountName {
						return true
					}
				}
				return false
			}
			if txn.Direction == txhelper.TransactionDirectionReceived {
				for _, input := range txn.Inputs {
					if isInArray(input.AccountName) {
						continue
					}
					accountNames = append(accountNames, input.AccountName)
				}
			} else {
				for _, output := range txn.Outputs {
					if isInArray(output.AccountName) {
						continue
					}
					accountNames = append(accountNames, output.AccountName)
				}
			}
			return strings.Join(accountNames, ", ")
		},
	}
}

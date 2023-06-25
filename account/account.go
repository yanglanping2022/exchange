package account

import (
	"github.com/urfave/cli/v2"
	"github.com/yanglanping2022/exchange/cex"
)

var AccountCommand = &cli.Command{
	Name:  "account",
	Usage: "list accounts or show account",
	Subcommands: []*cli.Command{
		{
			Name:   "list",
			Usage:  "list accounts",
			Action: listAccounts,
		},
		{
			Name:   "show",
			Usage:  "show accout",
			Action: showAccount,
		},
		{
			Name:   "balances",
			Usage:  "account balances",
			Action: showBalances,
		},
	},
}

func listAccounts(ctx *cli.Context) error {
	for _, cex := range cex.CexPool {
		cex.Name()
	}
	return nil
}

func showAccount(ctx *cli.Context) error {
	for _, cex := range cex.CexPool {
		cex.Name()
		cex.Account()
	}
	return nil
}

func showBalances(ctx *cli.Context) error {
	for _, cex := range cex.CexPool {
		cex.Name()
		cex.Balances()
	}
	return nil
}

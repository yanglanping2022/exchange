package cmd

import (
	"fmt"
	"log"

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
			Name:   "balances",
			Usage:  "account balances",
			Action: showBalances,
		},
		{
			Name:   "tradefee",
			Usage:  "trade fee",
			Action: showTradeFee,
		},
	},
}

func listAccounts(ctx *cli.Context) error {
	for _, ex := range cex.CexPool {
		log.Println(ex.Name())
	}
	return nil
}

func showBalances(ctx *cli.Context) error {
	for _, ex := range cex.CexPool {
		log.Println(ex.Name())

		balances, err := ex.Balances()
		if err != nil {
			log.Println(err)
			return err
		}

		for _, balance := range balances {
			log.Println(fmt.Sprintf("\t%s : %s", balance.Symbol, balance.Free))
		}

	}
	return nil
}

func showTradeFee(ctx *cli.Context) error {
	for _, ex := range cex.CexPool {
		log.Println(ex.Name())

		fee, err := ex.TradeFee(cex.ETHUSDT)
		if err != nil {
			log.Println(err)
			return err
		}

		log.Println(fmt.Sprintf("\tmaker commission: %f", fee.MakerCommission))
		log.Println(fmt.Sprintf("\ttaker commission: %f", fee.TakerCommission))
	}
	return nil
}

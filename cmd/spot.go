package cmd

import (
	"log"

	"github.com/urfave/cli/v2"
	"github.com/yanglanping2022/exchange/cex"
)

var SpotCommand = &cli.Command{
	Name:  "spot",
	Usage: "symbols list",
	Subcommands: []*cli.Command{
		{
			Name:   "symbols",
			Usage:  "symbols list",
			Action: listSymbols,
		},
	},
}

func listSymbols(ctx *cli.Context) error {
	for _, ex := range cex.CexPool {
		log.Println(ex.Name())

		symbols, err := ex.CurrencyPairs()
		if err != nil {
			log.Println(err)
			return err
		}

		log.Println(symbols)
	}
	return nil
}

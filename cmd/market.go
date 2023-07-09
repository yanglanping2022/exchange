package cmd

import (
	"log"

	"github.com/urfave/cli/v2"
	"github.com/yanglanping2022/exchange/cex"
)

var MarketCommand = &cli.Command{
	Name:  "market",
	Usage: "market info",
	Subcommands: []*cli.Command{
		{
			Name:   "bestorder",
			Usage:  "show bookticker",
			Action: bestOrder,
		},
	},
}

func bestOrder(ctx *cli.Context) error {
	for _, ex := range cex.CexPool {
		ex.Name()

		orderInfo, err := ex.BestOrder(cex.ETHUSDT)
		if err != nil {
			log.Println(err)
			return err
		}

		log.Println(orderInfo)
	}
	return nil
}

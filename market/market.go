package market

import (
	"github.com/urfave/cli/v2"
	"github.com/yanglanping2022/exchange/cex"
)

var MarketCommand = &cli.Command{
	Name:  "market",
	Usage: "market info",
	Subcommands: []*cli.Command{
		{
			Name:   "bookticker",
			Usage:  "show bookticker",
			Action: bookTicker,
		},
	},
}

func bookTicker(ctx *cli.Context) error {
	for _, cex := range cex.CexPool {
		cex.Name()
		cex.BookTicker()
	}
	return nil
}

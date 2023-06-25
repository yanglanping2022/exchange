package order

import (
	"github.com/urfave/cli/v2"
	"github.com/yanglanping2022/exchange/cex"
)

var OrderCommand = &cli.Command{
	Name:  "order",
	Usage: "order info",
	Subcommands: []*cli.Command{
		{
			Name:   "list",
			Usage:  "list orders",
			Action: listOrders,
		},
	},
}

func listOrders(ctx *cli.Context) error {
	for _, cex := range cex.CexPool {
		cex.Name()
		cex.AllOrders()
	}
	return nil
}

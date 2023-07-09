package cmd

import (
	"github.com/urfave/cli/v2"
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
	return nil
}

package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/yanglanping2022/exchange/account"
	"github.com/yanglanping2022/exchange/market"
	"github.com/yanglanping2022/exchange/order"
	"github.com/yanglanping2022/exchange/utils"
)

var app = &cli.App{
	Name:  "exchange",
	Usage: "exchange command tools",
	Flags: []cli.Flag{
		utils.EchoFlag,
	},
}

func init() {
	app.Commands = []*cli.Command{
		account.AccountCommand,
		market.MarketCommand,
		order.OrderCommand,
	}
}

func main() {
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

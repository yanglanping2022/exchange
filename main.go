package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/yanglanping2022/exchange/cex"
	"github.com/yanglanping2022/exchange/cmd"
	"github.com/yanglanping2022/exchange/config"
	"github.com/yanglanping2022/exchange/utils"
)

var app = &cli.App{
	Name:  "exchange",
	Usage: "exchange command tools",
	Flags: []cli.Flag{
		utils.ConfigFlag,
	},
}

func init() {
	app.Action = task
	app.Commands = []*cli.Command{
		cmd.AccountCommand,
		cmd.MarketCommand,
		cmd.OrderCommand,
		cmd.SpotCommand,
	}

	app.Before = func(ctx *cli.Context) error {
		file := ctx.String("config")
		// init config
		config.InitConf(file)
		// init cex
		cex.InitCex()
		return nil
	}

	app.After = func(ctx *cli.Context) error {
		return nil
	}
}

func main() {
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func task(ctx *cli.Context) error {
	log.Println("task start ......")
	for {
		chance(cex.CexPool[0], cex.CexPool[1], cex.BTCUSDT)
		chance(cex.CexPool[0], cex.CexPool[1], cex.ETHUSDT)
		chance(cex.CexPool[0], cex.CexPool[1], cex.FILUSDT)
	}
}

func chance(exchangeA cex.CEX, exchangeB cex.CEX, symbol int) {
	var (
		err        error
		bestOrderA *cex.BookOrderInfo
		bestOrderB *cex.BookOrderInfo
		tradeFeeA  *cex.TradeFeeInfo
		tradeFeeB  *cex.TradeFeeInfo
		bidCost    float32
		askCost    float32
	)

	if tradeFeeA, err = exchangeA.TradeFee(symbol); err != nil {
		log.Println("get exchangeA trade fee failed")
		return
	}

	if tradeFeeB, err = exchangeB.TradeFee(symbol); err != nil {
		log.Println("get exchangeB trade fee failed")
		return
	}

	if bestOrderA, err = exchangeA.BestOrder(symbol); err != nil {
		log.Println("get exchangeA best order failed")
		return
	}

	if bestOrderB, err = exchangeB.BestOrder(symbol); err != nil {
		log.Println("get exchangeB best order failed")
		return
	}

	// log.Println(symbol, bestOrderA, bestOrderB)

	//
	// chance:
	//    bidA - feeA - askB - feeB > 0
	//   (bidA - feeA) > (askB + feeB)
	//

	bidCost = bestOrderA.BidPrice - bestOrderA.BidPrice*tradeFeeA.TakerCommission
	askCost = bestOrderB.AskPrice + bestOrderB.AskPrice*tradeFeeB.TakerCommission
	if bidCost > askCost {
		log.Println(fmt.Sprintf("find chance, symbol: %d, profit: %f, bidA: %f, askB: %f",
			symbol, bidCost-askCost, bestOrderA.BidPrice, bestOrderB.AskPrice))
	}

	bidCost = bestOrderB.BidPrice - bestOrderB.BidPrice*tradeFeeB.TakerCommission
	askCost = bestOrderA.AskPrice + bestOrderA.AskPrice*tradeFeeA.TakerCommission
	if bidCost > askCost {
		log.Println(fmt.Sprintf("find change, symbol: %d, profit: %f, bidB: %f askA: %f",
			symbol, bidCost-askCost, bestOrderB.BidPrice, bestOrderA.AskPrice))
	}
}

package cex

import (
	"context"
	"errors"
	"strconv"

	binance_connector "github.com/binance/binance-connector-go"
	"github.com/yanglanping2022/exchange/config"
)

type BinanceCex struct {
	name   string
	client *binance_connector.Client
}

var binanceSymbolMap map[int]string = map[int]string{
	BTCUSDT: "BTCUSDT",
	ETHUSDT: "ETHUSDT",
}

func (r *BinanceCex) Name() string {
	return r.name
}

func (r *BinanceCex) NewClient() error {
	r.name = "Binance"

	r.client = binance_connector.NewClient(
		config.Conf.Binance.APIKey,
		config.Conf.Binance.SecretKey,
		config.Conf.Binance.BaseURL)

	return nil
}

func (r *BinanceCex) Balances() ([]BalanceInfo, error) {
	balances := []BalanceInfo{}

	resp, err := r.client.NewGetAccountService().
		Do(context.Background())
	if err != nil {
		return balances, err
	}

	for i := 0; i < len(resp.Balances); i++ {
		balance := resp.Balances[i]
		if balance.Asset == "USDT" || balance.Asset == "BNB" {
			balances = append(balances, BalanceInfo{Symbol: balance.Asset, Free: balance.Free})
		}
	}

	return balances, nil
}

func (r *BinanceCex) BestOrder(symbol int) (*BookOrderInfo, error) {
	symbolStr, exist := binanceSymbolMap[symbol]
	if !exist {
		return nil, errors.New("unknown symbol")
	}

	resp, err := r.client.NewTickerBookTickerService().
		Symbol(symbolStr).
		Do(context.Background())
	if err != nil {
		return nil, err
	}

	bookOrderInfo := BookOrderInfo{}

	bookOrderInfo.Name = r.name

	if bidPrice, err := strconv.ParseFloat(resp[0].BidPrice, 32); err != nil {
		return nil, err
	} else {
		bookOrderInfo.BidPrice = float32(bidPrice)
	}

	if bidQty, err := strconv.ParseFloat(resp[0].BidQty, 32); err != nil {
		return nil, err
	} else {
		bookOrderInfo.BidQty = float32(bidQty)
	}

	if askPrice, err := strconv.ParseFloat(resp[0].AskPrice, 32); err != nil {
		return nil, err
	} else {
		bookOrderInfo.AskPrice = float32(askPrice)
	}

	if askQty, err := strconv.ParseFloat(resp[0].AskQty, 32); err != nil {
		return nil, err
	} else {
		bookOrderInfo.AskQty = float32(askQty)
	}

	return &bookOrderInfo, nil
}

func (r *BinanceCex) TradeFee(symbol int) (*TradeFeeInfo, error) {
	// resp, err := r.client.NewTradeFeeService().
	// 	Symbol("ETHUSDT").
	// 	Do(context.Background())
	// if err != nil {
	// 	log.Println(err)
	// 	return err
	// }

	feeInfo := TradeFeeInfo{}
	feeInfo.MakerCommission = 0.001
	feeInfo.TakerCommission = 0.001
	return &feeInfo, nil
}

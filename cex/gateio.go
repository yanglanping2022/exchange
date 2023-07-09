package cex

import (
	"context"
	"errors"
	"log"
	"strconv"

	"github.com/antihax/optional"
	"github.com/gateio/gateapi-go/v6"
	"github.com/yanglanping2022/exchange/config"
)

type GateioCex struct {
	name    string
	client  *gateapi.APIClient
	context context.Context
}

var gateioSymbolMap map[int]string = map[int]string{
	BTCUSDT: "BTC_USDT",
	ETHUSDT: "ETH_USDT",
}

func (r *GateioCex) Name() string {
	return r.name
}

func (r *GateioCex) NewClient() error {
	r.name = "Gate.IO"

	r.client = gateapi.NewAPIClient(gateapi.NewConfiguration())
	r.context = context.WithValue(context.Background(),
		gateapi.ContextGateAPIV4,
		gateapi.GateAPIV4{
			Key:    config.Conf.Gateio.APIKey,
			Secret: config.Conf.Gateio.SecretKey,
		})

	return nil
}

func (r *GateioCex) Balances() ([]BalanceInfo, error) {
	balances := []BalanceInfo{}

	resp, _, err := r.client.WalletApi.GetTotalBalance(r.context, nil)
	if err != nil {
		return balances, err
	}

	balances = append(balances, BalanceInfo{Symbol: resp.Total.Currency, Free: resp.Total.Amount})

	return balances, nil
}

func (r *GateioCex) BestOrder(symbol int) (*BookOrderInfo, error) {
	symbolStr, exist := gateioSymbolMap[symbol]
	if !exist {
		return nil, errors.New("unknown symbol")
	}

	opts := gateapi.ListOrderBookOpts{}
	opts.Limit = optional.NewInt32(1)

	resp, _, err := r.client.SpotApi.ListOrderBook(r.context, symbolStr, &opts)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	orderBookInfo := BookOrderInfo{}

	orderBookInfo.Name = r.name

	if bidPrice, err := strconv.ParseFloat(resp.Bids[0][0], 32); err != nil {
		return nil, err
	} else {
		orderBookInfo.BidPrice = float32(bidPrice)
	}

	if bidQty, err := strconv.ParseFloat(resp.Bids[0][1], 32); err != nil {
		return nil, err
	} else {
		orderBookInfo.BidQty = float32(bidQty)
	}

	if askPrice, err := strconv.ParseFloat(resp.Asks[0][0], 32); err != nil {
		return nil, err
	} else {
		orderBookInfo.AskPrice = float32(askPrice)
	}

	if askQty, err := strconv.ParseFloat(resp.Asks[0][1], 32); err != nil {
		return nil, err
	} else {
		orderBookInfo.AskQty = float32(askQty)
	}

	return &orderBookInfo, nil
}

func (r *GateioCex) TradeFee(symbol int) (*TradeFeeInfo, error) {
	// opts := gateapi.GetTradeFeeOpts{}
	// opts.CurrencyPair = optional.NewString("ETH_USDT")
	// result, _, err := r.client.WalletApi.GetTradeFee(r.context, &opts)
	// if err != nil {
	// 	log.Println(err)
	// 	return err
	// }

	feeInfo := TradeFeeInfo{}
	feeInfo.MakerCommission = 0.002
	feeInfo.TakerCommission = 0.002
	return &feeInfo, nil
}

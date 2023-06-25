package cex

import (
	"context"
	"fmt"
	"log"

	binance_connector "github.com/binance/binance-connector-go"
)

type BinanceCex struct {
	name   string
	client *binance_connector.Client
}

func (r *BinanceCex) NewClient() error {
	r.name = "Binance"

	r.client = binance_connector.NewClient(
		"xxx",
		"yyy",
		"https://testnet.binance.vision")
	return nil
}

func (r *BinanceCex) Name() error {
	fmt.Println(r.name)
	return nil
}

func (r *BinanceCex) Account() error {
	resp, err := r.client.NewGetAccountService().
		Do(context.Background())
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Println(binance_connector.PrettyPrint(resp))
	return nil
}

func (r *BinanceCex) Balances() error {
	resp, err := r.client.NewGetAccountService().
		Do(context.Background())
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Println(binance_connector.PrettyPrint(resp.Balances))
	return nil
}

func (r *BinanceCex) BookTicker() error {
	resp, err := r.client.NewTickerBookTickerService().
		Symbols([]string{"ETHUSDT"}).
		// Symbol("ETHUSDT").
		Do(context.Background())
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Println(binance_connector.PrettyPrint(resp))
	return nil
}

func (r *BinanceCex) AllOrders() error {
	resp, err := r.client.NewGetAllOrdersService().
		Symbol("ETHUSDT").
		Do(context.Background())
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Println(binance_connector.PrettyPrint(resp))
	return nil
}

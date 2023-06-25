package cex

type CEX interface {
	Name() error
	Account() error
	NewClient() error
	Balances() error
	BookTicker() error
	AllOrders() error
}

var CexPool []CEX

func init() {
	binanceCex := BinanceCex{}
	binanceCex.NewClient()

	CexPool = append(CexPool, &binanceCex)
}

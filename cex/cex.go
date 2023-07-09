package cex

const (
	BTCUSDT = iota
	ETHUSDT
)

type BookOrderInfo struct {
	Name     string
	BidPrice float32
	AskPrice float32
	BidQty   float32
	AskQty   float32
}

type TradeFeeInfo struct {
	MakerCommission float32
	TakerCommission float32
}

type BalanceInfo struct {
	Symbol string
	Free   string
}

type CEX interface {
	Name() string
	NewClient() error
	Balances() ([]BalanceInfo, error)
	BestOrder(symbol int) (*BookOrderInfo, error)
	TradeFee(symbol int) (*TradeFeeInfo, error)
}

var CexPool []CEX

func InitCex() {
	// binance exchange
	binanceCex := BinanceCex{}
	binanceCex.NewClient()
	// gateio exchange
	gateioCex := GateioCex{}
	gateioCex.NewClient()

	CexPool = append(CexPool, &binanceCex)
	CexPool = append(CexPool, &gateioCex)
}

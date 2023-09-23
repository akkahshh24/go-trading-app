package trades

// model for the data received from binance
type Ticker struct {
	Symbol   string `json:"s"`
	Price    string `json:"p"`
	Quantity string `json:"q"`
	Time     int64  `json:"T"`
}

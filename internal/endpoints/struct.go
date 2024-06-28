package endpoints

import "time"

type DepthOrder struct {
	Price   float64 `json:"price"`
	BaseQty float64 `json:"baseQty"`
}

type Client struct {
	ClientName   string `json:"clientName"`
	ExchangeName string `json:"exchangeName"`
	Label        string `json:"label"`
	Pair         string `json:"pair"`
}

type HistoryOrder struct {
	ClientName          string    `json:"clientName"`
	ExchangeName        string    `json:"exchangeName"`
	Label               string    `json:"label"`
	Pair                string    `json:"pair"`
	Side                string    `json:"side"`
	Type                string    `json:"type"`
	BaseQty             float64   `json:"baseQty"`
	Price               float64   `json:"price"`
	AlgorithmNamePlaced string    `json:"algorithmNamePlaced"`
	LowestSellPrc       float64   `json:"lowestSellPrc"`
	HighestBuyPrc       float64   `json:"highestBuyPrc"`
	CommissionQuoteQty  float64   `json:"commissionQuoteQty"`
	TimePlaced          time.Time `json:"timePlaced"`
}

type GetOrderBookRequest struct {
	ExchangeName string `json:"exchange_name"`
	Pair         string `json:"pair"`
}

type GetOrderBookResponse struct {
	OrderBook []DepthOrder `json:"orderBook"`
	Error     error        `json:"error"`
}

type SaveOrderBookRequest struct {
	ExchangeName string       `json:"exchange_name"`
	Pair         string       `json:"pair"`
	OrderBook    []DepthOrder `json:"order_book"`
}

type SaveOrderBookResponse struct {
	Error error `json:"error"`
}

type GetOrderHistoryRequest struct {
	Client Client `json:"client"`
}

type GetOrderHistoryResponse struct {
	HistoryOrder []HistoryOrder `json:"historyOrder"`
	Error        error          `json:"error"`
}

type SaveOrderRequest struct {
	Client Client       `json:"client"`
	Order  HistoryOrder `json:"order"`
}

type SaveOrderResponse struct {
	Error error `json:"error"`
}

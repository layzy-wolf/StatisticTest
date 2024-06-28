package service

import (
	"time"
)

type App interface {
	GetOrderBook(exchangeName, pair string) ([]*DepthOrder, error)
	SaveOrderBook(exchangeName, pair string, orderBook []*DepthOrder) error
	GetOrderHistory(client *Client) ([]*HistoryOrder, error)
	SaveOrder(client *Client, order *HistoryOrder) error
}

type OrderHistory struct {
	ClientName          string
	ExchangeName        string
	Label               string
	Pair                string
	Side                string
	Type                string
	BaseQty             float64
	Price               float64
	AlgorithmNamePlaced string
	LowestSellPrc       float64
	HighestBuyPrc       float64
	CommissionQuoteQty  float64
	TimePlaced          time.Time
}

type OrderBook struct {
	Exchange string
	Pair     string
	Asks     []map[string]float64 `gorm:"serializer:json"`
	Bids     []map[string]float64 `gorm:"serializer:json"`
}

type DepthOrder struct {
	Price   float64
	BaseQty float64
}

type HistoryOrder struct {
	ClientName          string
	ExchangeName        string
	Label               string
	Pair                string
	Side                string
	Type                string
	BaseQty             float64
	Price               float64
	AlgorithmNamePlaced string
	LowestSellPrc       float64
	HighestBuyPrc       float64
	CommissionQuoteQty  float64
	TimePlaced          time.Time
}

type Client struct {
	ClientName   string
	ExchangeName string
	Label        string
	Pair         string
}

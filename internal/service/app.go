package service

import (
	"gorm.io/gorm"
	"time"
)

type Service struct {
	store *gorm.DB
}

func NewStatisticService(store *gorm.DB) *Service {
	return &Service{store: store}
}

// GetOrderBook select orderBook by exchangeName and pair and returns Asks and error
func (s *Service) GetOrderBook(exchangeName, pair string) ([]*DepthOrder, error) {
	var orderBook []OrderBook
	var orders []*DepthOrder
	err := s.store.Model(OrderBook{}).Where(OrderBook{Exchange: exchangeName, Pair: pair}).Find(&orderBook).Error

	if err != nil {
		return nil, err
	}
	for _, order := range orderBook {
		for _, val := range order.Asks {
			orders = append(orders, &DepthOrder{
				Price:   val["price"],
				BaseQty: val["base_qty"],
			})
		}
	}

	return orders, err
}

// SaveOrderBook inserts new value to orderBook
func (s *Service) SaveOrderBook(exchangeName, pair string, orderBook []*DepthOrder) error {
	var orders []map[string]float64

	for _, val := range orderBook {
		orders = append(orders, map[string]float64{"price": val.Price, "base_qty": val.BaseQty})
	}

	order := OrderBook{
		Exchange: exchangeName,
		Pair:     pair,
		Asks:     orders,
		Bids:     orders,
	}

	err := s.store.Create(&order).Error

	if err != nil {
		return err
	}

	return nil
}

// GetOrderHistory select orderHistory by client and returns all records
func (s *Service) GetOrderHistory(client *Client) ([]*HistoryOrder, error) {
	var historyOrder []*HistoryOrder
	var orderHistory []OrderHistory

	tx := s.store.Model(OrderHistory{}).Where(OrderHistory{ClientName: client.ClientName}).Find(&orderHistory)

	if tx.Error != nil {
		return nil, tx.Error
	}

	for _, val := range orderHistory {
		historyOrder = append(historyOrder, &HistoryOrder{
			ClientName:          val.ClientName,
			ExchangeName:        val.ExchangeName,
			Label:               val.Label,
			Pair:                val.Pair,
			Side:                val.Side,
			Type:                val.Type,
			BaseQty:             val.BaseQty,
			Price:               val.Price,
			AlgorithmNamePlaced: val.AlgorithmNamePlaced,
			LowestSellPrc:       val.LowestSellPrc,
			HighestBuyPrc:       val.HighestBuyPrc,
			CommissionQuoteQty:  val.CommissionQuoteQty,
			TimePlaced:          val.TimePlaced,
		})
	}

	return historyOrder, nil
}

// SaveOrder inserts client and historyOrder into new record
func (s *Service) SaveOrder(client *Client, order *HistoryOrder) error {
	t, _ := time.Parse(time.DateTime, order.TimePlaced.String())

	ord := OrderHistory{
		ClientName:          client.ClientName,
		ExchangeName:        client.ExchangeName,
		Label:               client.Label,
		Pair:                client.Pair,
		Side:                order.Side,
		Type:                order.Type,
		BaseQty:             order.BaseQty,
		Price:               order.Price,
		AlgorithmNamePlaced: order.AlgorithmNamePlaced,
		LowestSellPrc:       order.LowestSellPrc,
		HighestBuyPrc:       order.HighestBuyPrc,
		CommissionQuoteQty:  order.CommissionQuoteQty,
		TimePlaced:          t,
	}

	tx := s.store.Create(&ord)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

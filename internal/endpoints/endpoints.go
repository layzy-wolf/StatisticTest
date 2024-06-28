package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/layzy-wolf/StatisticTest/internal/service"
)

type Endpoints struct {
	GetOrderBook    endpoint.Endpoint
	SaveOrderBook   endpoint.Endpoint
	GetOrderHistory endpoint.Endpoint
	SaveOrder       endpoint.Endpoint
}

// NewEndpoints wrap all service logic to go-kit endpoints
func NewEndpoints(srv *service.Service) Endpoints {
	return Endpoints{
		GetOrderBook:    MakeGetOrderBook(srv),
		SaveOrderBook:   MakeSaveOrderBook(srv),
		GetOrderHistory: MakeGetOrderHistory(srv),
		SaveOrder:       MakeSaveOrder(srv),
	}
}

func MakeGetOrderBook(srv *service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		var orders []DepthOrder
		req := request.(GetOrderBookRequest)
		success, err := srv.GetOrderBook(req.ExchangeName, req.Pair)

		if err != nil {
			return GetOrderBookResponse{
				OrderBook: nil,
				Error:     err,
			}, err
		}

		for _, val := range success {
			orders = append(orders, DepthOrder{
				Price:   val.Price,
				BaseQty: val.BaseQty,
			})
		}

		return GetOrderBookResponse{
			OrderBook: orders,
			Error:     nil,
		}, nil
	}
}

func MakeSaveOrderBook(srv *service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		var orders []*service.DepthOrder
		req := request.(SaveOrderBookRequest)

		for _, val := range req.OrderBook {
			orders = append(orders, &service.DepthOrder{
				Price:   val.Price,
				BaseQty: val.BaseQty,
			})
		}
		err = srv.SaveOrderBook(req.ExchangeName, req.Pair, orders)

		if err != nil {
			return SaveOrderBookResponse{Error: err}, err
		}
		return SaveOrderBookResponse{Error: nil}, nil
	}
}

func MakeGetOrderHistory(srv *service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		var history []HistoryOrder
		req := request.(GetOrderHistoryRequest)

		success, err := srv.GetOrderHistory(&service.Client{
			ClientName:   req.Client.ClientName,
			ExchangeName: req.Client.ExchangeName,
			Label:        req.Client.Label,
			Pair:         req.Client.Pair,
		})

		if err != nil {
			return GetOrderHistoryResponse{
				HistoryOrder: nil,
				Error:        err,
			}, err
		}

		for _, val := range success {
			history = append(history, HistoryOrder{
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

		return GetOrderHistoryResponse{
			HistoryOrder: history,
			Error:        nil,
		}, err
	}
}

func MakeSaveOrder(srv *service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(SaveOrderRequest)
		err = srv.SaveOrder(&service.Client{
			ClientName:   req.Client.ClientName,
			ExchangeName: req.Client.ExchangeName,
			Label:        req.Client.Label,
			Pair:         req.Client.Pair,
		}, &service.HistoryOrder{
			ClientName:          req.Order.ClientName,
			ExchangeName:        req.Order.ExchangeName,
			Label:               req.Order.Label,
			Pair:                req.Order.Pair,
			Side:                req.Order.Side,
			Type:                req.Order.Type,
			BaseQty:             req.Order.BaseQty,
			Price:               req.Order.Price,
			AlgorithmNamePlaced: req.Order.AlgorithmNamePlaced,
			LowestSellPrc:       req.Order.LowestSellPrc,
			HighestBuyPrc:       req.Order.HighestBuyPrc,
			CommissionQuoteQty:  req.Order.CommissionQuoteQty,
			TimePlaced:          req.Order.TimePlaced,
		})
		if err != nil {
			return SaveOrderResponse{Error: err}, err
		}
		return SaveOrderResponse{Error: nil}, nil
	}
}

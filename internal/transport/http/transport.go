package transport

import (
	"context"
	"encoding/json"
	kitHttp "github.com/go-kit/kit/transport/http"
	"github.com/layzy-wolf/StatisticTest/internal/endpoints"
	"github.com/layzy-wolf/StatisticTest/internal/service"
	"gorm.io/gorm"
	"net/http"
)

type AppServer struct {
	GetOrderBook    *kitHttp.Server
	SaveOrderBook   *kitHttp.Server
	GetOrderHistory *kitHttp.Server
	SaveOrder       *kitHttp.Server
}

func decodeGetOrderBookRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request endpoints.GetOrderBookRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeSaveOrderBookRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request endpoints.SaveOrderBookRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeGetOrderHistoryRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request endpoints.GetOrderHistoryRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeSaveOrderRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request endpoints.SaveOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// MakeAppHandler wraps all endpoints to http.Server handlers
func MakeAppHandler(store *gorm.DB) *AppServer {
	srv := service.NewStatisticService(store)
	en := endpoints.NewEndpoints(srv)
	return &AppServer{
		GetOrderBook:    kitHttp.NewServer(en.GetOrderBook, decodeGetOrderBookRequest, encodeResponse),
		SaveOrderBook:   kitHttp.NewServer(en.SaveOrderBook, decodeSaveOrderBookRequest, encodeResponse),
		GetOrderHistory: kitHttp.NewServer(en.GetOrderHistory, decodeGetOrderHistoryRequest, encodeResponse),
		SaveOrder:       kitHttp.NewServer(en.SaveOrder, decodeSaveOrderRequest, encodeResponse),
	}
}

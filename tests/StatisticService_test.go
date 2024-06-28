package tests

import (
	"github.com/layzy-wolf/StatisticTest/config"
	"github.com/layzy-wolf/StatisticTest/internal/service"
	"github.com/layzy-wolf/StatisticTest/internal/store"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	client = service.Client{
		ClientName:   "test",
		ExchangeName: "test",
		Label:        "test",
		Pair:         "test",
	}
	order = service.HistoryOrder{
		ClientName:          "test",
		ExchangeName:        "test",
		Label:               "test",
		Pair:                "test",
		Side:                "test",
		Type:                "test",
		BaseQty:             10,
		Price:               10,
		AlgorithmNamePlaced: "test",
		LowestSellPrc:       10,
		HighestBuyPrc:       10,
		CommissionQuoteQty:  10,
		TimePlaced:          time.Now().UTC(),
	}
	exchangeName = "test"
	pair         = "test"
	orderBook    = []*service.DepthOrder{
		{Price: 10, BaseQty: 10},
		{Price: 10, BaseQty: 10},
	}
)

func TestSaveOrderBook(t *testing.T) {
	s := store.NewStore(config.Cfg{
		Port:       8080,
		DBSocket:   "localhost:9000",
		DBName:     "default",
		DBUser:     "default",
		DBPassword: "",
	})

	srv := service.NewStatisticService(s)

	err := srv.SaveOrderBook(exchangeName, pair, orderBook)

	assert.Nil(t, err)
}

func TestSaveOrder(t *testing.T) {
	s := store.NewStore(config.Cfg{
		Port:       8080,
		DBSocket:   "localhost:9000",
		DBName:     "default",
		DBUser:     "default",
		DBPassword: "",
	})

	srv := service.NewStatisticService(s)

	err := srv.SaveOrder(&client, &order)

	assert.Nil(t, err)
}

func TestGetOrderHistory(t *testing.T) {
	s := store.NewStore(config.Cfg{
		Port:       8080,
		DBSocket:   "localhost:9000",
		DBName:     "default",
		DBUser:     "default",
		DBPassword: "",
	})

	srv := service.NewStatisticService(s)

	history, err := srv.GetOrderHistory(&client)

	assert.Nil(t, err)
	assert.NotEmpty(t, history)
}

func TestGetOrderBook(t *testing.T) {
	s := store.NewStore(config.Cfg{
		Port:       8080,
		DBSocket:   "localhost:9000",
		DBName:     "default",
		DBUser:     "default",
		DBPassword: "",
	})

	srv := service.NewStatisticService(s)

	orders, err := srv.GetOrderBook(exchangeName, pair)

	assert.Nil(t, err)
	assert.NotEmpty(t, orders)
}

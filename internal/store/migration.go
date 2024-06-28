package store

import (
	"github.com/layzy-wolf/StatisticTest/internal/service"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		service.DepthOrder{},
		service.OrderBook{},
		service.OrderHistory{},
	)
	return err
}

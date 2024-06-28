package store

import (
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/layzy-wolf/StatisticTest/config"
	"github.com/layzy-wolf/StatisticTest/internal/service"
	gormClickhouse "gorm.io/driver/clickhouse"
	"gorm.io/gorm"
	"log"
	"time"
)

func NewStore(cfg config.Cfg) *gorm.DB {
	var (
		conn = clickhouse.OpenDB(&clickhouse.Options{
			Addr: []string{cfg.DBSocket},
			Auth: clickhouse.Auth{
				Database: fmt.Sprintf("%v", cfg.DBName),
				Username: fmt.Sprintf("%v", cfg.DBUser),
				Password: fmt.Sprintf("%v", cfg.DBPassword),
			},
			Settings: clickhouse.Settings{
				"max_execution_time": 60,
			},
			DialTimeout: 5 * time.Second,
			Compression: &clickhouse.Compression{
				Method: clickhouse.CompressionLZ4,
			},
			Debug: true,
		})
	)

	if err := conn.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			log.Panicf("Exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
	}

	db, err := gorm.Open(gormClickhouse.New(gormClickhouse.Config{Conn: conn}))

	if err != nil {
		log.Panicln(err)
	}

	has := db.Migrator().HasTable(service.DepthOrder{}) &&
		db.Migrator().HasTable(service.OrderHistory{}) &&
		db.Migrator().HasTable(service.OrderBook{})

	if !has {
		if err := Migrate(db); err != nil {
			log.Fatalln(err)
		}
	}

	return db
}

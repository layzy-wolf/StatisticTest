package main

import (
	"context"
	"fmt"
	"github.com/layzy-wolf/StatisticTest/config"
	"github.com/layzy-wolf/StatisticTest/internal/store"
	transport "github.com/layzy-wolf/StatisticTest/internal/transport/http"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

// main call layers and realize graceful shutdown
func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg := config.Load()
	s := store.NewStore(cfg)
	r := transport.HTTPHandler(s)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", cfg.Port),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()
	stop()
	log.Println("Server shutting down gracefully, press Ctrl+C to force")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}

package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Hymiside/test-task-appmagic/pkg/cache"
	"github.com/Hymiside/test-task-appmagic/pkg/config"
	"github.com/Hymiside/test-task-appmagic/pkg/handler"
	"github.com/Hymiside/test-task-appmagic/pkg/server"
	"github.com/Hymiside/test-task-appmagic/pkg/service"
)

func main() {
	cfgServ := config.InitConfig()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	srv := &server.Server{}
	h := &handler.Handler{}

	ch := cache.NewCache()
	services := service.NewService(*ch)

	if err := services.SetInfoGas(); err != nil {
		log.Panicf("error set info: %v", err)
	}

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
		select {
		case <-quit:
			cancel()
		case <-ctx.Done():
			return
		}
	}()

	if err := srv.RunServer(ctx, h.InitHandler(*services), cfgServ); err != nil {
		log.Panicf("failed to run server: %v", err)
	}
	log.Printf("server was stopped")
}

package main

import (
	"context"
	"github.com/Hymiside/test-task-appmagic/pkg/repository"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Hymiside/test-task-appmagic/pkg/config"
	"github.com/Hymiside/test-task-appmagic/pkg/handler"
	"github.com/Hymiside/test-task-appmagic/pkg/server"
	"github.com/Hymiside/test-task-appmagic/pkg/service"
)

func main() {
	cfgServ, cfgRepo := config.InitConfig()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	srv := &server.Server{}
	h := &handler.Handler{}

	repo, err := repository.NewRepository(ctx, cfgRepo)
	if err != nil {
		log.Panicf("falied to create redis reository:%v", err)
	}
	services := service.NewService(repo)

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

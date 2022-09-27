package config

import (
	"github.com/Hymiside/test-task-appmagic/pkg/repository"
	"os"

	"github.com/Hymiside/test-task-appmagic/pkg/server"
	"github.com/joho/godotenv"
)

// InitConfig инициализирует конфиг
func InitConfig() (server.ConfigServer, repository.ConfigRepository) {
	_ = godotenv.Load()
	host, _ := os.LookupEnv("SERVER_HOST")
	port, _ := os.LookupEnv("SERVER_PORT")
	hostRepo, _ := os.LookupEnv("REPO_HOST")
	portRepo, _ := os.LookupEnv("REPO_PORT")

	cfgServ := server.ConfigServer{
		Host: host,
		Port: port,
	}
	cfgRepo := repository.ConfigRepository{
		Host: hostRepo,
		Port: portRepo,
	}

	return cfgServ, cfgRepo
}

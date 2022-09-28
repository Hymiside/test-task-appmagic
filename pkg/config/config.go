package config

import (
	"os"

	"github.com/Hymiside/test-task-appmagic/pkg/server"
	"github.com/joho/godotenv"
)

// InitConfig инициализирует конфиг
func InitConfig() server.ConfigServer {
	_ = godotenv.Load()
	host, _ := os.LookupEnv("SERVER_HOST")
	port, _ := os.LookupEnv("SERVER_PORT")

	return server.ConfigServer{
		Host: host,
		Port: port,
	}
}

//+build !test

package main

import (
	"fmt"
	"go-scrum-poker-bot/config"
	"go-scrum-poker-bot/storage"
	"go-scrum-poker-bot/ui"
	"go-scrum-poker-bot/web/clients"
	clients_middleware "go-scrum-poker-bot/web/clients/middleware"
	"go-scrum-poker-bot/web/server"
	server_middleware "go-scrum-poker-bot/web/server/middleware"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	logger := log.New(os.Stdout, "INFO: ", log.LstdFlags)
	config := config.NewConfig()
	redisCLI := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
		DB:   config.Redis.DB,
	})
	userStorage := storage.NewUserRedisStorage(redisCLI)
	sessionStorage := storage.NewSessionRedisStorage(redisCLI)
	builder := ui.NewBuilder(config)
	webClient := clients.NewBasicClient(
		&http.Client{
			Timeout: 5 * time.Second,
		},
		[]clients.Middleware{
			clients_middleware.Auth(config.Slack.Token),
			clients_middleware.JsonContentType,
			clients_middleware.Log(logger),
		},
	)

	app := server.NewServer(
		logger,
		webClient,
		builder,
		userStorage,
		sessionStorage,
		[]server.Middleware{server_middleware.Recover(logger), server_middleware.Log(logger), server_middleware.Json},
	)
	app.Serve(config.App.ServerAddress)
}

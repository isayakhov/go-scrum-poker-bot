package server

import (
	"context"
	"go-scrum-poker-bot/storage"
	"go-scrum-poker-bot/ui"
	"go-scrum-poker-bot/web/clients"
	"go-scrum-poker-bot/web/server/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"time"
)

type Middleware func(next http.Handler) http.Handler

type Server struct {
	healthy        int32
	middleware     []Middleware
	logger         *log.Logger
	webClient      clients.Client
	uiBuilder      *ui.Builder
	userStorage    storage.UserStorage
	sessionStorage storage.SessionStorage
}

func NewServer(
	logger *log.Logger,
	webClient clients.Client,
	uiBuilder *ui.Builder,
	userStorage storage.UserStorage,
	sessionStorage storage.SessionStorage,
	middleware []Middleware,
) *Server {
	return &Server{
		logger:         logger,
		webClient:      webClient,
		uiBuilder:      uiBuilder,
		userStorage:    userStorage,
		sessionStorage: sessionStorage,
		middleware:     middleware,
	}
}

func (s *Server) setupRouter() http.Handler {
	router := http.NewServeMux()
	router.Handle(
		"/healthcheck",
		handlers.Healthcheck(),
	)
	router.Handle(
		"/play-poker",
		handlers.PlayPokerCommand(s.webClient, s.uiBuilder),
	)
	router.Handle(
		"/interactivity",
		handlers.InteractionCallback(s.userStorage, s.sessionStorage, s.uiBuilder, s.webClient),
	)

	return router
}

func (s *Server) setupMiddleware(router http.Handler) http.Handler {
	handler := router
	for _, middleware := range s.middleware {
		handler = middleware(handler)
	}

	return handler
}

func (s *Server) Serve(address string) {
	server := &http.Server{
		Addr:         address,
		Handler:      s.setupMiddleware(s.setupRouter()),
		ErrorLog:     s.logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		s.logger.Println("Server is shutting down...")
		atomic.StoreInt32(&s.healthy, 0)

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			s.logger.Fatalf("Could not gracefully shutdown the server: %v\n", err)
		}
		close(done)
	}()

	s.logger.Println("Server is ready to handle requests at", address)
	atomic.StoreInt32(&s.healthy, 1)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		s.logger.Fatalf("Could not listen on %s: %v\n", address, err)
	}

	<-done
	s.logger.Println("Server stopped")
}

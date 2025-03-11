package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/meahmadhassan/go-restapi/internal/config"
	"github.com/meahmadhassan/go-restapi/internal/http/handlers/student"
)

func main() {
	// load config
	cfg := config.MustLoad()
	// custom log if needed

	// database setup

	// route setup
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New())

	// setup server http
	server := http.Server {
		Addr: cfg.Addr,
		Handler: router,
	}

	slog.Info("Server started", slog.String("address", cfg.Addr))

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start server")
		}
	}()

	<- done 
	
	slog.Info("Shutting down the server")
	
	ctx, cancel := context.WithTimeout(context.Background(), 5* time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("Server shutdown successfully!")
 
}



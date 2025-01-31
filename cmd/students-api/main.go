package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Raghavendra/students-api/internal/config"
	student "github.com/Raghavendra/students-api/internal/config/http/handlers"
	"github.com/Raghavendra/students-api/internal/storage/sqlite"
)

func main() {

	//load config

	cfg := config.MustLoad()

	//database setup

	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("storage initilaized", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))

	//setup router

	router := http.NewServeMux()

	router.HandleFunc("GET /api/students", student.GetStudentsById(storage))
	router.HandleFunc("GET/api/students/{id}", student.GetById(storage))
	router.HandleFunc("GET/api/students", student.GetList(storage))

	//setup server(HTTP)

	server := http.Server{

		Addr:    cfg.Addr,
		Handler: router,
	}

	slog.Info("starting server", slog.String("address", cfg.Addr))
	fmt.Println("server started", cfg.HTTPServer.Addr)

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("failed to start server")
		}
	}()
	<-done

	//for structured logging
	slog.Info("Shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		slog.Error("failed to shutdown server")
	}

	slog.Info("server shutdown")
}

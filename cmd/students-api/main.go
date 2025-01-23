package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Raghavendra/students-api/internal/config"
)

func main() {

	//load config

	cfg := config.MustLoad()

	//database setup
	//setup router

	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to students api"))
	})

	//setup server(HTTP)

	server := http.Server{

		Addr:    cfg.Addr,
		Handler: router,
	}
	fmt.Printf("server started", cfg.HTTPServer.Addr)

	done := make(chan os.Signal, 1)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("failed to start server")
		}
	}()
	<-done
}

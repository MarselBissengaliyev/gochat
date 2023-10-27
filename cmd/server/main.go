package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"net/http"

	"github.com/MarselBissengaliyev/gochat/pkg/handler"
	"github.com/MarselBissengaliyev/gochat/pkg/service"
)

var (
	port = flag.String("port", "8080", "port used for ws connection")
)

func main() {
	flag.Parse()

	mux := http.NewServeMux()
	services := service.NewService(mux)
	handlers := handler.NewHandler(services)

	srv := handlers.InitHandlers(*port)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("listen and serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print("Graccefully shutting down")
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("error shuttind down server: %v", err)
	}
}

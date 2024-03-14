package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v3/log"
)

func main() {
	srv := http.Server{
		Addr: "localhost:8080",
	}

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
		<-sigint

		if err := srv.Shutdown(context.Background()); err != nil {
			log.Error("shutdown error", err)
		}
		close(idleConnsClosed)
	}()

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Error("run error", err)
		}
	}()
}

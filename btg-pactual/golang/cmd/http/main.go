package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/buemura/btg-challenge/config"
	"github.com/buemura/btg-challenge/infra/database"
	"github.com/buemura/btg-challenge/infra/queue"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	config.LoadEnv()
	database.Connect()
	queue.StartConsume()
}

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Secure())

	host := ":" + config.HTTP_PORT

	go func() {
		if err := e.Start(host); err != nil && http.ErrServerClosed != err {
			panic(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, os.Interrupt, syscall.SIGINT)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	fmt.Println("Stopping...")

	if err := e.Shutdown(ctx); err != nil {
		panic(err)
	}
	fmt.Println("Server stopped")
}

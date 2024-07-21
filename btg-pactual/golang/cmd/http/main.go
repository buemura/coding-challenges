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
	"github.com/buemura/btg-challenge/internal/infra/database"
	"github.com/buemura/btg-challenge/internal/infra/queue"
	"github.com/buemura/btg-challenge/internal/modules/order"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	config.LoadEnv()
	database.Connect()

}

func main() {
	go queue.StartConsume()

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	order.SetupRoutes(e)

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

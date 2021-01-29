package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"dominos.com/orders"
	"dominos.com/orders/models"
	"dominos.com/orders/server"
	"dominos.com/orders/services"
	"github.com/go-kit/kit/log"
	"github.com/joho/godotenv"
)

func main() {
	var (
		httpAddr = flag.String("port", ":8080", "HTTP listen address")
	)
	flag.Parse()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	err := godotenv.Load(os.ExpandEnv("../.env"))
	if err != nil {
		panic("missing os environment vars.")
	}

	conn := orders.NewConnection(os.Getenv("ORDERS_DB_USERNAME"), os.Getenv("ORDERS_DB_PASSWORD"), os.Getenv("ORDERS_DB_NAME"), logger)
	conn.DB.AutoMigrate(&models.Order{})
	conn.DB.AutoMigrate(&models.OrderItem{})

	orderRepository := services.NewOrderRepository(conn)
	tlogger := orders.NewLogService(logger)
	var orderService services.OrderService
	{
		orderService = services.NewOrderService(orderRepository, logger)
		orderService = services.NewLoggingOrderServiceMiddleware(logger, tlogger)(orderService)
	}
	orderItemRepository := services.NewOrderItemRepository(conn)
	var orderItemService services.OrderItemService
	{
		orderItemService = services.NewOrderItemService(orderItemRepository, logger)
	}
	httpHandler := server.MakeHTTPHandler(orderService, orderItemService, logger)

	errors := make(chan error)
	go func() {
		osSignal := make(chan os.Signal)
		signal.Notify(osSignal, syscall.SIGINT, syscall.SIGTERM)
		errors <- fmt.Errorf("%s", <-osSignal)
	}()

	go func() {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		errors <- http.ListenAndServe(*httpAddr, httpHandler)
	}()

	logger.Log("exit", <-errors)
}

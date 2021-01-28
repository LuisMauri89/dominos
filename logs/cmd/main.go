package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/dominos/logs"
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

	conn := logs.NewConnection(os.Getenv("LOGS_DB_USERNAME"), os.Getenv("LOGS_DB_PASSWORD"), os.Getenv("LOGS_DB_NAME"), logger)
	defer conn.DB.Close()
	tlogsepository := logs.NewTraceLogRepository(conn)
	var tlogsService logs.TraceLogService
	{
		tlogsService = logs.NewTraceLogService(tlogsepository, logger)
		tlogsService = logs.NewLoggingTraceLogServiceMiddleware(logger)(tlogsService)
	}
	httpHandler := logs.MakeHTTPHandler(tlogsService, logger)

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

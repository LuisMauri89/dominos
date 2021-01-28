package test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/dominos/logs"
	"github.com/go-kit/kit/log"
	"github.com/gofrs/uuid"
	"github.com/joho/godotenv"
)

const createTableTlogsQuery = `CREATE TABLE IF NOT EXISTS tlogs(id UUID PRIMARY KEY, timestamp TEXT NOT NULL, service_name TEXT NOT NULL, caller TEXT, event TEXT NOT NULL, extra TEXT)`

var conn logs.Connection
var logger log.Logger
var tlogsepository logs.TraceLogRepository
var tlogsService logs.TraceLogService

func TestMain(m *testing.M) {
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	err := godotenv.Load(os.ExpandEnv("../.env"))
	if err != nil {
		logger.Log("env error", "load env vars failed.")
	}

	conn = logs.NewConnection(os.Getenv("LogsDbUsername"), os.Getenv("LogsDbPassword"), os.Getenv("LogsDbName"), logger)
	defer conn.DB.Close()
	tlogsepository = logs.NewTraceLogRepository(conn)
	tlogsService = logs.NewTraceLogService(tlogsepository, logger)
	tlogsService = logs.NewLoggingTraceLogServiceMiddleware(logger)(tlogsService)

	checkTableTlogsExist()
	test := m.Run()
	ClearTableTlogs()

	os.Exit(test)
}

func checkTableTlogsExist() error {
	if _, err := conn.DB.Exec(createTableTlogsQuery); err != nil {
		logger.Log("error", err)
		return err
	}
	logger.Log("database", "connected.")
	return nil
}

func ClearTableTlogs() {
	conn.DB.Exec("DELETE FROM tlogs")
	logger.Log("database", "table tlogs cleaned.")
}

func SeedTableTlogsOneTlog() (logs.TraceLog, error) {
	tlog := logs.TraceLog{
		ServiceName: "ORDERS",
		Caller:      "Create Order Method",
		Event:       "NEW",
		Extra:       "Pizza x 1, CocaCola x 2. $19.99",
	}

	uuid, _ := uuid.NewV4()
	id := uuid.String()
	tlog.ID = id

	now := time.Now()
	timestamp := now.Unix()
	tlog.TimeStamp = strconv.FormatInt(timestamp, 10)

	err := conn.DB.QueryRow("INSERT INTO tlogs(id, timestamp, service_name, caller, event, extra) VALUES($1, $2, $3, $4, $5, $6) RETURNING id",
		tlog.ID,
		tlog.TimeStamp,
		tlog.ServiceName,
		tlog.Caller,
		tlog.Event,
		tlog.Extra).Scan(&tlog.ID)

	if err != nil {
		return logs.TraceLog{}, err
	}

	return tlog, nil
}

func ExecuteRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	httpHandler := logs.MakeHTTPHandler(tlogsService, logger)

	httpHandler.ServeHTTP(rr, req)

	return rr
}

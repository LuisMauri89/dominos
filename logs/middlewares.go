package logs

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
)

type LoggingTraceLogServiceMiddleware func(s TraceLogService) TraceLogService

type loggingTraceLogServiceMiddleware struct {
	TraceLogService
	logger log.Logger
}

func NewLoggingTraceLogServiceMiddleware(logger log.Logger) LoggingTraceLogServiceMiddleware {
	return func(next TraceLogService) TraceLogService {
		return &loggingTraceLogServiceMiddleware{next, logger}
	}
}

func (mw *loggingTraceLogServiceMiddleware) FindAll(ctx context.Context) ([]TraceLog, error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "FindAllTlogs", "took", time.Since(begin))
	}(time.Now())
	return mw.TraceLogService.FindAll(ctx)
}

func (mw *loggingTraceLogServiceMiddleware) Create(ctx context.Context, tlog TraceLog) error {
	defer func(begin time.Time) {
		mw.logger.Log("method", "CreateTlog", "took", time.Since(begin))
	}(time.Now())
	return mw.TraceLogService.Create(ctx, tlog)
}

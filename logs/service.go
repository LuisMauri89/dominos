package logs

import "context"

// TraceLogService - implement this interface to inject service with functional operations over TraceLog entity.
type TraceLogService interface {
	FindAll(ctx context.Context) ([]TraceLog, error)
	Create(ctx context.Context, tlog TraceLog) error
}

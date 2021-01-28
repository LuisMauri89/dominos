package logs

import "context"

// TraceLogRepository - handles quering to postgresql logs database.
type TraceLogRepository interface {
	FindAll(ctx context.Context) ([]TraceLog, error)
	Create(ctx context.Context, tlog TraceLog) error
}

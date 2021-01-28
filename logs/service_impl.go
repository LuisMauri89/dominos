package logs

import (
	"context"
	"strconv"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/gofrs/uuid"
)

type traceLogService struct {
	repository TraceLogRepository
	logger     log.Logger
}

// NewTraceLogService - returns new instance implementation of TraceLogService interface.
func NewTraceLogService(repository TraceLogRepository, logger log.Logger) TraceLogService {
	return &traceLogService{
		repository: repository,
		logger:     logger,
	}
}

func (s *traceLogService) FindAll(ctx context.Context) ([]TraceLog, error) {
	logs, err := s.repository.FindAll(ctx)
	if err != nil {
		s.logger.Log("error:", err)
		return nil, err
	}
	s.logger.Log("findall:", "success")
	return logs, nil
}

func (s *traceLogService) Create(ctx context.Context, tlog TraceLog) error {
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	tlog.ID = id

	now := time.Now()
	timestamp := now.Unix()
	tlog.TimeStamp = strconv.FormatInt(timestamp, 10)

	if err := s.repository.Create(ctx, tlog); err != nil {
		s.logger.Log("error:", err)
		return err
	}
	s.logger.Log("create:", "success")
	return nil
}

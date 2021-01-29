package deliveries

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/gofrs/uuid"
)

type deliveryService struct {
	repository DeliveryRepository
	logger     log.Logger
	tlogger    LogsService
}

func NewDeliveryService(repository DeliveryRepository, logger log.Logger, tlogger LogsService) DeliveryService {
	return &deliveryService{
		repository: repository,
		logger:     logger,
		tlogger:    tlogger,
	}
}

func (s *deliveryService) FindAll(ctx context.Context) ([]Delivery, error) {
	deliveries, err := s.repository.FindAll(ctx)
	if err != nil {
		s.logger.Log("error:", err)
		return nil, err
	}
	s.logger.Log("findall:", "success")
	return deliveries, nil
}

func (s *deliveryService) Create(ctx context.Context, delivery Delivery) error {
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	delivery.ID = id

	if err := s.repository.Create(ctx, delivery); err != nil {
		s.logger.Log("error:", err)
		return err
	}
	s.logger.Log("create:", "success")

	go s.tlogger.SaveLog(TLog{
		ServiceName: "DELIVERIES",
		Caller:      "Delivery->Create",
		Event:       "POST",
		Extra:       "Create new delivery.",
	})
	return nil
}

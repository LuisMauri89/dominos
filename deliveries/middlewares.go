package deliveries

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
)

type LoggingDeliveryServiceMiddleware func(s DeliveryService) DeliveryService

type loggingDeliveryServiceMiddleware struct {
	DeliveryService
	logger log.Logger
}

func NewLoggingDeliveryServiceMiddleware(logger log.Logger) LoggingDeliveryServiceMiddleware {
	return func(next DeliveryService) DeliveryService {
		return &loggingDeliveryServiceMiddleware{next, logger}
	}
}

func (mw *loggingDeliveryServiceMiddleware) FindAll(ctx context.Context) ([]Delivery, error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "FindAllTDely", "took", time.Since(begin))
	}(time.Now())
	return mw.DeliveryService.FindAll(ctx)
}

func (mw *loggingDeliveryServiceMiddleware) Create(ctx context.Context, td Delivery) error {
	defer func(begin time.Time) {
		mw.logger.Log("method", "CreateTDely", "took", time.Since(begin))
	}(time.Now())
	return mw.DeliveryService.Create(ctx, td)
}
func (mw *loggingDeliveryServiceMiddleware) Delete(ctx context.Context, td Delivery) error {
	defer func(begin time.Time) {
		mw.logger.Log("method", "DeleteTDely", "took", time.Since(begin))
	}(time.Now())
	return mw.DeliveryService.Delete(ctx, td)
}
func (mw *loggingDeliveryServiceMiddleware) Update(ctx context.Context, td Delivery) error {
	defer func(begin time.Time) {
		mw.logger.Log("method", "UpdateTDely", "took", time.Since(begin))
	}(time.Now())
	return mw.DeliveryService.Update(ctx, td)
}

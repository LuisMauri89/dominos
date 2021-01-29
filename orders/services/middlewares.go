package services

import (
	"context"
	"encoding/json"
	"time"

	"dominos.com/orders"
	"dominos.com/orders/models"
	"github.com/go-kit/kit/log"
)

type LoggingOrderServiceMiddleware func(s OrderService) OrderService

type loggingOrderServiceMiddleware struct {
	OrderService
	tlogger orders.LogsService
	logger  log.Logger
}

func NewLoggingOrderServiceMiddleware(logger log.Logger, tlogger orders.LogsService) LoggingOrderServiceMiddleware {
	return func(next OrderService) OrderService {
		return &loggingOrderServiceMiddleware{next, tlogger, logger}
	}
}

func (mw *loggingOrderServiceMiddleware) FindAll(ctx context.Context) ([]models.Order, error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "FindAllOrders", "took", time.Since(begin))
	}(time.Now())

	collection, err := mw.OrderService.FindAll(ctx)
	tlog := orders.TLog{
		ServiceName: "ORDERS",
		Caller:      "Order->FindAll",
		Event:       "GET",
		Extra:       "Find all orders.",
	}

	if err != nil {
		extra, _ := json.Marshal(err)
		tlog.Extra = string(extra)
	}
	go mw.tlogger.SaveLog(tlog)

	return collection, err
}

func (mw *loggingOrderServiceMiddleware) GetByID(ctx context.Context, id string) (models.Order, error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "GetByID", "took", time.Since(begin))
	}(time.Now())

	order, err := mw.OrderService.GetByID(ctx, id)
	tlog := orders.TLog{
		ServiceName: "ORDERS",
		Caller:      "Order->GetByID",
		Event:       "GET",
		Extra:       "Get order by ID.",
	}

	if err != nil {
		extra, _ := json.Marshal(err)
		tlog.Extra = string(extra)
	}
	go mw.tlogger.SaveLog(tlog)

	return order, err
}

func (mw *loggingOrderServiceMiddleware) Create(ctx context.Context, order models.Order) error {
	defer func(begin time.Time) {
		mw.logger.Log("method", "Create", "took", time.Since(begin))
	}(time.Now())

	err := mw.OrderService.Create(ctx, order)
	tlog := orders.TLog{
		ServiceName: "ORDERS",
		Caller:      "Order->Create",
		Event:       "POST",
		Extra:       "Create new order.",
	}

	if err != nil {
		extra, _ := json.Marshal(err)
		tlog.Extra = string(extra)
	}
	go mw.tlogger.SaveLog(tlog)

	return err
}

func (mw *loggingOrderServiceMiddleware) Delete(ctx context.Context, id string) error {
	defer func(begin time.Time) {
		mw.logger.Log("method", "Delete", "took", time.Since(begin))
		mw.tlogger.SaveLog(orders.TLog{
			ServiceName: "ORDERS",
			Caller:      "Order->Delete",
			Event:       "DELETE",
			Extra:       "Delete order by ID.",
		})
	}(time.Now())

	err := mw.OrderService.Delete(ctx, id)
	tlog := orders.TLog{
		ServiceName: "ORDERS",
		Caller:      "Order->Delete",
		Event:       "DELETE",
		Extra:       "Delete order by ID.",
	}

	if err != nil {
		extra, _ := json.Marshal(err)
		tlog.Extra = string(extra)
	}
	go mw.tlogger.SaveLog(tlog)

	return err
}

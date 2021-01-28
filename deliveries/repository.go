package DeliveriesV2

import "context"

// TraceLogRepository - handles quering to postgresql logs database.
type TraceLogRepository interface {
	FindAll(ctx context.Context) ([]Delivery, error)
	Create(ctx context.Context, td Delivery) error
	Delete(ctx context.Context, td Delivery) error
	Update(ctx context.Context, td Delivery) error
}

package deliveries

import "context"

type DeliveryRepository interface {
	FindAll(ctx context.Context) ([]Delivery, error)
	Create(ctx context.Context, td Delivery) error
	GetByStatus(ctx context.Context, status string) (models.De}, error)
}

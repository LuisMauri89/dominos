package deliveries

import "context"

type DeliveryRepository interface {
	FindAll(ctx context.Context) ([]Delivery, error)
	Create(ctx context.Context, td Delivery) error
	Delete(ctx context.Context, td Delivery) error
	Update(ctx context.Context, td Delivery) error
}

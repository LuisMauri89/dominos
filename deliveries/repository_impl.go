package deliveries

import (
	"context"
	"errors"
	"sync"
)

var (
	ErrInconsistentIDs = errors.New("inconsistent IDs")
	ErrAlreadyExists   = errors.New("already exists")
	ErrNotFound        = errors.New("not found")
)

type deliveryRepository struct {
	mtx  sync.RWMutex
	conn Connection
}

func NewDeliveryRepository(conn Connection) DeliveryRepository {
	return &deliveryRepository{
		conn: conn,
	}
}

func (r *deliveryRepository) FindAll(ctx context.Context) ([]Delivery, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	rows, err := r.conn.DB.Query("SELECT id, order_id, status, name, final_price, address, description FROM deliveries")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	deliveries := []Delivery{}

	for rows.Next() {
		var d Delivery
		if err := rows.Scan(&d.ID, &d.OrderID, &d.Status, &d.Name, &d.FinalPrice, &d.Address, &d.Description); err != nil {
			return nil, err
		}
		deliveries = append(deliveries, d)
	}

	return deliveries, nil
}

func (r *deliveryRepository) Create(ctx context.Context, delivery Delivery) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	err := r.conn.DB.QueryRow("INSERT INTO deliveries(id, order_id, status, name, final_price, address, description) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		delivery.ID,
		delivery.OrderID,
		delivery.Status,
		delivery.Name,
		delivery.FinalPrice,
		delivery.Address,
		delivery.Description).Scan(&delivery.ID)

	if err != nil {
		return err
	}

	return nil
}
func (r *deliveryRepository) GetByStatus(ctx context.Context) ([]Delivery, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	rows, err := r.conn.DB.Query("SELECT id, order_id, status, name, final_price, address, description FROM deliveries WHERE status=PENDING")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	deliveries := []Delivery{}

	for rows.Next() {
		var d Delivery
		if err := rows.Scan(&d.ID, &d.OrderID, &d.Status, &d.Name, &d.FinalPrice, &d.Address, &d.Description); err != nil {
			return nil, err
		}
		deliveries = append(deliveries, d)
	}

	return deliveries, nil
}

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

type DeliveryRepository struct {
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
	rows, err := r.conn.DB.Query("SELECT id, orderId, status, to, finalPrice,address,description FROM tDely")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tDely := []Delivery{}

	for rows.Next() {
		var td Delivery
		if err := rows.Scan(&td.ID, &td.OrderId, &td.Status, &td.To, &td.FinalPrice, &td.Address, &td.Description); err != nil {
			return nil, err
		}
		tDely = append(tDely, td)
	}

	return tDely, nil
}

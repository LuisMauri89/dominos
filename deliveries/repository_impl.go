package DeliveriesV2

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

type traceLogRepository struct {
	mtx  sync.RWMutex
	conn Connection
}

func NewTraceLogRepository(conn Connection) TraceLogRepository {
	return &traceLogRepository{
		conn: conn,
	}
}

func (r *traceLogRepository) FindAll(ctx context.Context) ([]Delivery, error) {
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

func (r *traceLogRepository) Create(ctx context.Context, td Delivery) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	err := r.conn.DB.QueryRow("INSERT INTO tDely(id,orderId, status, to, finalPrice, address, description) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id",
	td.ID,
	td.OrderID,
	td.Status,
	td.To,
	td.FinalPrice,
	td.Address,
	td.Description

	if err != nil {
		return err
	}

	return nil
}
func (r *traceLogRepository) Delete(ctx context.Context, td Delivery) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	err := r.conn.DB.QueryRow("DELETE FROM tDely WHERE id = id",
	td.ID,
		td.OrderID,
		td.Status,
		td.To,
		td.FinalPrice,
		td.Address,
		td.Description


	if err != nil {
		return err
	}

	return nil
}

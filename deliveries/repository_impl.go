package deliveries

import (
	"context"
	"errors"
	"sync"
)

var (
	// ErrInconsistentIDs - update entity and retrieved entity are incompatible.
	ErrInconsistentIDs = errors.New("inconsistent IDs")
	// ErrAlreadyExists - create / update entity already exist.
	ErrAlreadyExists = errors.New("already exists")
	//ErrNotFound - desired entity can not be found.
	ErrNotFound = errors.New("not found")
)

type traceLogRepository struct {
	mtx  sync.RWMutex
	conn Connection
}

// NewTraceLogRepository - returns new instance implementation of TraceLogRepository interface.
func NewTraceLogRepository(conn Connection) TraceLogRepository {
	return &traceLogRepository{
		conn: conn,
	}
}

func (r *traceLogRepository) ListarTodo(ctx context.Context) ([]TraceLog, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	rows, err := r.conn.DB.Query("SELECT id, orderId, status, to, finalPrice,address,description, extra FROM tdeliveries")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tdeliveries := []TraceLog{}

	for rows.Next() {
		var tdeli TraceLog
		if err := rows.Scan(&tdeli.ID, &tdeli.TimeStamp, &tdeli.ServiceName, &tdeli.Caller, &tdeli.Event, &tdeli.Extra); err != nil {
			return nil, err
		}
		tdeliveries = append(tdeliveries, tdeli)
	}

	return tdeliveries, nil
}

func (r *traceLogRepository) Create(ctx context.Context, tdeli TraceLog) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	err := r.conn.DB.QueryRow("INSERT INTO tdeliveries(id,orderId, status, to, finalPrice, address, description) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		tdeli.ID,
		tdeli.OrderID,
		tdeli.Status,
		tdeli.To,
		tdeli.FinalPrice,
		tdeli.Address,
		tdeli.Description,


	if err != nil {
		return err
	}

	return nil
}
func (r *traceLogRepository) Delete(ctx context.Context, tdeli TraceLog) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	err := r.conn.DB.QueryRow("DELETE FROM tdeliveries WHERE id = id",
		tdeli.ID,
		tdeli.OrderID,
		tdeli.Status,
		tdeli.To,
		tdeli.FinalPrice,
		tdeli.Address,
		tdeli.Description,


	if err != nil {
		return err
	}

	return nil
}
func (r *traceLogRepository) Update(ctx context.Context, tdeli TraceLog) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	err := r.conn.DB.QueryRow("UPDATE tdeliveries SET orderId = "" , status = "", + descripcion ="", + To ="",+ finalPrice ="", + address =""+ description ="" , WHERE id = id")
		tdeli.ID,
		tdeli.OrderID,
		tdeli.Status,
		tdeli.To,
		tdeli.FinalPrice,
		tdeli.Address,
		tdeli.Description,


	if err != nil {
		return err
	}

	return nil
}

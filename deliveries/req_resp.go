package deliveries

type FindAllRequest struct {
}

type FindAllResponse struct {
	TDely []Delivery `json:"tdely"`
	Err   error      `json:"error,omitempty"`
}

func (r FindAllResponse) error() error { return r.Err }

type CreateOrderRequest struct {
	Delivery models.Delivery `json:"deliveries"`
}
type CreateOrderResponse struct {
	Err error `json:"error,omitempty"`
}

func (r CreateOrderResponse) error() error { return r.Err }

type GetByStatusRequest struct {
	Status string `json:"status"`
}
type GetByStatusResponse struct {
	Delivery models.Delivery `json:"deliveries"`
	Err      error           `json:"error,omitempty"`
}

func (r GetByStatusResponse) error() error { return r.Err }

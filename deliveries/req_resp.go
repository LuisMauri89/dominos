package deliveries

type FindAllRequest struct {
}

type FindAllResponse struct {
	TDely []Delivery `json:"tdely"`
	Err   error      `json:"error,omitempty"`
}

func (r FindAllResponse) error() error { return r.Err }

type CreateOrderRequest struct {
	Order models.Order `json:"order"`
}
type CreateOrderResponse struct {
	Err error `json:"error,omitempty"`
}

func (r CreateOrderResponse) error() error { return r.Err }

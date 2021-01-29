package deliveries

type FindAllRequest struct {
}

type FindAllResponse struct {
	TDely []Delivery `json:"tdely"`
	Err   error      `json:"error,omitempty"`
}

func (r FindAllResponse) error() error { return r.Err }

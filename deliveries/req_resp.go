package DeliveriesV2

type FindAllRequest struct {
}

type CreateRequest struct {
	TDely TraceLog `json:"td"`
}

type FindAllResponse struct {
	TDely []Delivery `json:"tdely"`
	Err   error      `json:"error,omitempty"`
}

func (r FindAllResponse) error() error { return r.Err }

type CreateResponse struct {
	Err error `json:"error,omitempty"`
}

func (r CreateResponse) error() error { return r.Err }

type DeleteResponse struct {
	Err error `json:"error,omitempty"`
}

func (r DeleteResponse) error() error { return r.Err }

type UpdateResponse struct {
	Err error `json:"error,omitempty"`
}

func (r UpdateResponse) error() error { return r.Err }

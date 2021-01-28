package advent

type ListRequest struct {
}

type UpdateRequest struct {
	ID     string
	Advert Delivery
}

type DeleteRequest struct {
	ID string
}

type GetByIdRequest struct {
	ID string
}
type GetByStatus struct {
	Status string
}

type ListResponse struct {
	Ads map[string]Delivery `json:"ads"`
	Err error               `json:"error,omitempty"`
}

func (r ListResponse) error() error { return r.Err }

type UpdateResponse struct {
	Err error `json:"error,omitempty"`
}

func (r UpdateResponse) error() error { return r.Err }

type DeleteResponse struct {
	Err error `json:"error,omitempty"`
}

func (r DeleteResponse) error() error { return r.Err }

type GetByIdResponse struct {
	Advert Delivery `json:"advert"`
	Err    error    `json:"error,omitempty"`
}

func (r GetByIdResponse) error() error { return r.Err }

type GetByStatusResponse struct {
	Advert Delivery `json:"advert"`
	Err    error    `json:"error,omitempty"`
}

func (r GetByStatusResponse) error() error { return r.Err }

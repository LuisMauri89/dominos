package logs

// FindAllRequest - maps GET logs request object.
type FindAllRequest struct {
}

// CreateRequest - maps POST create log request object.
type CreateRequest struct {
	Tlog TraceLog `json:"tlog"`
}

// FindAllResponse - maps GET logs response object.
type FindAllResponse struct {
	Tlogs []TraceLog `json:"tlogs"`
	Err   error      `json:"error,omitempty"`
}

func (r FindAllResponse) error() error { return r.Err }

// CreateResponse - maps POST create log response object.
type CreateResponse struct {
	Err error `json:"error,omitempty"`
}

func (r CreateResponse) error() error { return r.Err }

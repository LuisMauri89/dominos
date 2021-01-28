package logs

// TraceLog - entity that holds a log's data.
type TraceLog struct {
	ID          string `json:"id"`
	TimeStamp   string `json:"timestamp"`
	ServiceName string `json:"serviceName"`
	Caller      string `json:"caller,omitempty"`
	Event       string `json:"event"`
	Extra       string `json:"extra,omitempty"`
}

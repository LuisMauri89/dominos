package deliveries

type Delivery struct {
	ID          string  `json:"id"`
	OrderID     string  `json:"orderid"`
	Status      string  `json:"status"` // PENDING - READY - DELIVERED
	Name        string  `json:"name"`
	FinalPrice  float64 `json:"finalprice"`
	Address     string  `json:"address"`
	Description string  `json:"description,omitempty"`
}

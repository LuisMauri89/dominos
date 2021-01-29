package models

type Delivery struct {
	ID          string `json:"id"`
	OrderID     string `json:"orderid,omitempty"`
	Status      string `json:"status,omitempty"` // PENDING - READY - DELIVERED
	To          string `json:"to,omitempty"`
	FinalPrice  int    `json:"finalprice,omitempty"`
	Address     string `json:"address,omitempty"`
	Description string `json:"description,omitempty"`
}

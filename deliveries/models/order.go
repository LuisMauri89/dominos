package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/gofrs/uuid"
)

type Order struct {
	ID          string      `gorm:"primary_key" json:"id"`
	OrderID     string      `gorm:"size:255;not null" json:"orderid"`
	Status      string      `gorm:"size:255;not null;" json:"status"`
	To          string      `gorm:"size:255;not null;" json:"to"`
	FinalPrice  float64     `json:"finalprice"`
	Address     string      `gorm:"size:255;not null;" json:"address"`
	Description string      `gorm:"size:255;not null;" json:"description"`
	CreatedAt   time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeldatedAt  time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"deldated_at"`
	Items       []OrderItem `gorm:"foreignKey:OrderID" json:"items"`
}

func (o *Order) Prepare() {
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	o.ID = id

	o.OrderID = html.EscapeString(strings.TrimSpace(o.OrderID))
	o.Status = html.EscapeString(strings.TrimSpace(o.Status))
	o.To = html.EscapeString(strings.TrimSpace(o.To))
	o.Address = html.EscapeString(strings.TrimSpace(o.Address))
	o.Description = html.EscapeString(strings.TrimSpace(o.Description))
	o.CreatedAt = time.Now()
	o.UpdatedAt = time.Now()
}

func (o *Order) Validate() error {
	if o.OrderID == "" {
		return errors.New("required OrderID")
	}
	if o.Status == "" {
		return errors.New("required insert status")
	}
	if o.To == "" {
		return errors.New("required name to")
	}
	if o.Description == "" {
		return errors.New("required order description")
	}

	if o.Address == "" {
		return errors.New("required address")
	}
	if o.FinalPrice <= 0.0 {
		return errors.New("total price should be other than cero")
	}

	return nil
}

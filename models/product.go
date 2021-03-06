package models

import (
	"time"
)

// Product struct
type Product struct {
	tableName struct{}  `sql:"products"`
	ID        string    `json:"id" pg:",pk"`
	Name      string    `json:"name"`
	Expiry    time.Time `json:"expiry"`
	OrderID   string    `json:"order_id"`
	Delete    bool      `json:"delete" pg:",use_zero"`
}

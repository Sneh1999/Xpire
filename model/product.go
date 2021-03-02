package model

import (
	"time"
)

// Product struct
type Product struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Expiry time.Time `json:"expiry"`
}
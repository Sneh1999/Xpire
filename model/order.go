package model

import (
	"time"
)

// Order struct
type Order struct {
	ID   int64 	  `json:"id"`
	Date *time.Time `json:"date"`
}

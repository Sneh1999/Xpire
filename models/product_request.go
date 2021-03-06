package models

type ProductRequest struct {
	Name    string `json:"name"`
	Expiry  string `json:"expiry"`
	OrderID string `json:"order_id"`
}

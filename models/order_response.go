package models

type OrderResponse struct {
	Message string `json:"message"`
	OrderID string `json:"order_id"`
}

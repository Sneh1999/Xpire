package models

type CreateProductResponse struct {
	Message   string `json:"message"`
	ProductID string `json:"id"`
}

type GetProductResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Expiry  string `json:"expiry"`
	OrderID string `json:"order_id"`
}

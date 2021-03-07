package models

type CreateProductResponse struct {
	Message   string `json:"message"`
	ProductID string `json:"product_id"`
}

type GetProductResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Expiry  string `json:"expiry"`
	OrderID string `json:"order_id"`
}

type EditProductResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Expiry  string `json:"expiry"`
	OrderID string `json:"order_id"`
}

type DeleteProductResponse struct {
	Message string `json:"message"`
}

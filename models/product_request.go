package models

type CreateProductRequest struct {
	Name    string `json:"name"`
	Expiry  string `json:"expiry"`
	OrderID string `json:"order_id"`
}

type GetProductRequest struct {
	OrderID   string `json:"order_id"`
	ProductID string `json:"product_id"`
}

type EditProductRequest struct {
	Name      string `json:"name"`
	Expiry    string `json:"expiry"`
	OrderID   string `json:"order_id"`
	ProductID string `json:"product_id"`
}

type DeleteProductRequest struct {
	OrderID   string `json:"order_id"`
	ProductID string `json:"product_id"`
}

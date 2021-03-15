package models

type GetOrderRequest struct {
	ID string `json:"id"`
}

type CreateOrderRequest struct {
	Name string `json:"name"`
}

type EditOrderRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type DeleteOrderRequest struct {
	ID string `json:"id"`
}

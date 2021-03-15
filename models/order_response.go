package models

type CreateOrderResponse struct {
	Message string `json:"message"`
	ID      string `json:"order_id"`
}

type GetOrderResponse struct {
	ID       string     `json:"order_id"`
	UserID   string     `json:"user_id"`
	Products []*Product `json:"products"`
}

type EditOrderResponse struct {
	ID       string     `json:"order_id"`
	UserID   string     `json:"user_id"`
	Products []*Product `json:"products"`
}

type DeleteOrderResponse struct {
	Message string `json:"message"`
}

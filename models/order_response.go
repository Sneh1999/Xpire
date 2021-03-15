package models

type CreateOrderResponse struct {
	Message string `json:"message"`
	ID      string `json:"id"`
}

type GetOrderResponse struct {
	ID       string     `json:"id"`
	UserID   string     `json:"user_id"`
	Products []*Product `json:"products"`
}

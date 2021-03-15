package models

type GetOrderRequest struct {
	OrderId string `json:"order_id"`
}

type EditOrderRequest struct {
	ID       string     `json:"order_id"`
	Products []*Product `json:"products"`
}

type DeleteOrderRequest struct {
	ID string `json:"order_id"`
}

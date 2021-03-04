package models

type OrderRequest struct {
	Products []Product `json:"products"`
	UserID   string    `json:"userId"`
}

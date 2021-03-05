package models

// User struct
type User struct {
	tableName struct{} `sql:"user"`
	ID        string   `json:"id" pg:",pk"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Password  string   `json:"password"`
	Orders    []*Order `json:"orders" pg:"rel:has-many"`
}

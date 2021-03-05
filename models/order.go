package models

type Order struct {
	tableName struct{} `sql:"orders"`
	ID        string   `json:"id"  pg:",pk"`
	UserID    string
	Products  []*Product `json:"products" pg:"rel:has-many"`
}

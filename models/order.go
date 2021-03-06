package models

type Order struct {
	tableName struct{}   `sql:"orders"`
	ID        string     `json:"id"  pg:",pk"`
	Name      string     `json:"name"`
	UserID    string     `json:"user_id"`
	Products  []*Product `json:"products" pg:"rel:has-many`
	Delete    bool       `json:"delete" pg:",use_zero"`
}

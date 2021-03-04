package models

// User struct
type User struct {
	ID       string `json:"id" pg:",pk`
	Name     string `required:"true" json:"name"`
	Email    string `required:"true" json:"email"`
	Password string `required:"true" json:"password"`
}

package model

// User struct
type User struct {
	ID     int64    `json:"id" pg:"default:gen_random_uuid(),pk"`
	Name   string   `required: "true" json:"name"`
	Email  string `required: "true" json:"email"`
	Password string `required: "true" json:"password"`
}

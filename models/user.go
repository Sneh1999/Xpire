package models

import uuid "github.com/satori/go.uuid"

// User struct
type User struct {
	ID       uuid.UUID `json:"id" pg:",pk`
	Name     string `required: "true" json:"name"`
	Email    string `required: "true" json:"email"`
	Password string `required: "true" json:"password"`
}

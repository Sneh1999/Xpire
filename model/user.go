package model

// User struct
type User struct {
    ID     int64 `json:"id"`
    Name   string `json:"name"`
    Emails []string `json:"emails"`
}
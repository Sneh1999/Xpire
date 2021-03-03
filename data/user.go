package data

import "github.com/Sneh1999/Xpire/model"

//AddUser helps in adding new user to the database
func (db* DatabaseService) AddUser(user* model.User) error{
	_, err := db.DB.Model(user).Insert()
	return err
}
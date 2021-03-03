package data

import "github.com/Sneh1999/Xpire/models"

//AddUser helps in adding new user to the database
func (db *DatabaseService) AddUser(user *models.User) error {
	_, err := db.DB.Model(user).Insert()
	return err
}

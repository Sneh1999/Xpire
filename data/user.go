package data

import "github.com/Sneh1999/Xpire/models"

//GetUser
func (db *DatabaseService) GetUser(user *models.User) error {
	err := db.db.Model(user).Where("email = ?", user.Email).Select()
	return err
}

//AddUser helps in adding new user to the database
func (db *DatabaseService) AddUser(user *models.User) error {
	_, err := db.db.Model(user).Insert()
	return err
}

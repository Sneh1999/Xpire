package data

import "github.com/Sneh1999/Xpire/models"

//GetOrder helps us getting an order
func (db *DatabaseService) GetOrder(order *models.Order) error {
	err := db.db.Model(order).Where("ID = ?", order.ID).Select()
	return err
}

//AddOrder helps in adding new order to the database
func (db *DatabaseService) AddOrder(order *models.Order) error {
	_, err := db.db.Model(order).Insert()
	return err
}

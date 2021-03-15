package data

import (
	"github.com/Sneh1999/Xpire/models"
	"github.com/go-pg/pg/v10/orm"
)

//GetOrder helps us getting an order
func (db *DatabaseService) GetOrder(order *models.Order) error {
	err := db.db.Model(order).Relation("Products", func(q *orm.Query) (*orm.Query, error) {
		return q.Where("order_id = ? and delete IS FALSE", order.ID), nil
	}).Where("ID = ?", order.ID).Select()
	return err
}

//AddOrder helps in adding new order to the database
func (db *DatabaseService) AddOrder(order *models.Order) error {
	_, err := db.db.Model(order).Insert()
	return err
}

// EditOrder: helps in editing an existing order
func (db *DatabaseService) EditOrder(order *models.Order) error {
	_, err := db.db.Model(order).Where("id = ?", order.ID).Update()
	return err
}

//DeleteOrder helps us editing a product
func (db *DatabaseService) DeleteOrder(order *models.Order) error {
	_, err := db.db.Model(order).Where("id = ?", order.ID).Update()
	return err
}

package data

import "github.com/Sneh1999/Xpire/models"

//GetProduct helps us getting a product
func (db *DatabaseService) GetProduct(product *models.Product) error {
	err := db.db.Model(product).Where("id = ?", product.ID).Select()
	return err
}

//AddProduct helps in adding new product to the database
func (db *DatabaseService) AddProduct(product *models.Product) error {
	_, err := db.db.Model(product).Insert()
	return err
}

//EditProduct helps us editing a product
func (db *DatabaseService) EditProduct(product *models.Product) error {
	_, err := db.db.Model(product).Where("id = ?", product.ID).Update()
	return err
}

//DeleteProduct helps us in marking a product as deleted
func (db *DatabaseService) DeleteProduct(product *models.Product) error {
	_, err := db.db.Model(product).Where("id = ?", product.ID).Update()
	return err
}

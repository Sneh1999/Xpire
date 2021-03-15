package data

import (
	"github.com/Sneh1999/Xpire/models"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/sirupsen/logrus"
)

type DatabaseService struct {
	db  *pg.DB
	log *logrus.Logger
}

// NewDatabaseService is used to initialize the database
func NewDatabaseService(databaseConfig *models.DatabaseConfig, log *logrus.Logger) (*DatabaseService, error) {

	db := pg.Connect(&pg.Options{
		Addr:     databaseConfig.DBAddr,
		User:     databaseConfig.DBUser,
		Password: databaseConfig.DBPassword,
		Database: databaseConfig.DBName,
	})

	databaseService := &DatabaseService{
		db:  db,
		log: log,
	}

	err := databaseService.createSchema()

	if err != nil {
		log.Errorf("An error occured while creating schema : %s", err)
		return nil, err
	}

	return databaseService, nil
}

func (d *DatabaseService) createSchema() error {
	models := []interface{}{
		(*models.User)(nil),
		(*models.Product)(nil),
		(*models.Order)(nil),
	}

	for _, model := range models {
		err := d.db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp: false,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

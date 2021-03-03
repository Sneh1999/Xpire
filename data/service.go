package data

import (
	"github.com/Sneh1999/Xpire/model"
	"github.com/go-pg/pg"
	"github.com/sirupsen/logrus"
)

type DatabaseService struct {
	DB  *pg.DB
	log *logrus.Logger
}

// NewDatabaseService is used to initialize the database
func NewDatabaseService(databaseConfig *model.DatabaseConfig, log *logrus.Logger) (*DatabaseService, error) {

	db := pg.Connect(&pg.Options{
		Addr:     databaseConfig.DBAddr,
		User:     databaseConfig.DBUser,
		Password: databaseConfig.DBPassword,
		Database: databaseConfig.DBName,
	})

	databaseService := &DatabaseService{
		DB:  db,
		log: log,
	}

	err := databaseService.createSchema()

	if err != nil {
		log.Errorf(" An error occured while creating schema : %s", err)
		return nil, err
	}

	return databaseService, nil
}

func (d *DatabaseService) createSchema() error {
	models := []interface{}{
		(*model.User)(nil),
		(*model.Product)(nil),
		(*model.Order)(nil),
	}

	for _, model := range models {
		err := d.DB.Model(model).CreateTable(nil)
		if err != nil {
			return err
		}
	}
	return nil
}

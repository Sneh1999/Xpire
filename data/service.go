package data

import (
	"github.com/Sneh1999/Xpire/model"
	"github.com/go-pg/pg"
	"github.com/sirupsen/logrus"
)

type DatabaseService struct {
	DB *pg.DB
	log *logrus.Logger
}


func NewDatabaseService(databaseConfig *model.DatabaseConfig,log *logrus.Logger) *DatabaseService {
	db := pg.Connect(&pg.Options{
		Addr:     databaseConfig.DBAddr,
		User:     databaseConfig.DBUser,
		Password: databaseConfig.DBPassword,
		Database: databaseConfig.DBName,
	})
	log.WithField("address", databaseConfig.DBAddr).Info("Database connected on address")
	
	return &DatabaseService{
		DB: db,
		log: log,
	}
}
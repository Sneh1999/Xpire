package data

import (
	"database/sql"

	"github.com/Sneh1999/Xpire/model"
	"github.com/go-pg/pg"
)

type DatabaseService struct {
	DB *sql.DB
}

func NewDatabaseService(databaseConfig *model.DatabaseConfig) *DatabaseService {
	db := pg.Connect(&pg.Options{
		Addr:     databaseConfig.Address,
		User:     databaseConfig.User,
		Password: databaseConfig.Password,
		Database: databaseConfig.Dbname,
	})
	return &DatabaseService{
		DB: db,
	}
}
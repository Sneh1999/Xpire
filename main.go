package main

import (
	"os"

	"github.com/Sneh1999/Xpire/model"
)

func main() {

	databaseConfig := &model.DatabaseConfig{
		Address: os.Getenv("APP_DB_ADDR"),
		User: os.Getenv("APP_DB_USERNAME"),
		Password: os.Getenv("APP_DB_PASSWORD"),
		Dbname : os.Getenv("APP_DB_NAME"),
	}
	
	app := NewApp(databaseConfig);

	app.Run(":8000")
}
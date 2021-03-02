package main

import (
	"github.com/Sneh1999/Xpire/data"
	"github.com/Sneh1999/Xpire/model"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// App struct to initialize the Router and the Database
type App struct {
	Router *mux.Router
	dbService *data.DatabaseService
}

// Initialize the app 
func Initialize(databaseConfig *model.DatabaseConfig) (*App) {

	databaseService := data.NewDatabaseService(databaseConfig)
	return &App{
		dbService: databaseService,
	}
}

// Run the app
func(a *App) Run(addr string) {

 }

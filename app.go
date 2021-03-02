package main

import (
	"net/http"

	"github.com/Sneh1999/Xpire/data"
	"github.com/Sneh1999/Xpire/model"
	"github.com/Sneh1999/Xpire/router"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

// App struct to initialize the Router and the Database
type App struct {
	routerService *router.RouterService
	dbService *data.DatabaseService
	log *logrus.Logger
}

// Initialize the app 
func Initialize(databaseConfig *model.DatabaseConfig, log *logrus.Logger) (*App) {

	databaseService := data.NewDatabaseService(databaseConfig,log)
	routerService := router.NewRouterService(log)
	return &App{
		dbService: databaseService,
		routerService: routerService,
		log: log,
	}
	
}

// Run the app
func(a *App) Run(addr string) {
	//TODO: add error handling - server crashing 
	http.ListenAndServe(addr, a.routerService.Router)
 }

package main

import (
	"net/http"

	"github.com/Sneh1999/Xpire/data"
	"github.com/Sneh1999/Xpire/models"
	"github.com/Sneh1999/Xpire/router"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

// App struct to initialize the Router and the Database
type App struct {
	routerService *router.RouterService
	dbService     *data.DatabaseService
	log           *logrus.Logger
}

// Initialize the app
func Initialize(config *models.Config, log *logrus.Logger) *App {

	databaseService, err := data.NewDatabaseService(&config.DatabaseConfig, log)
	if err != nil {
		log.WithError(err).Fatal("Error setting up database service")
	}

	log.WithField("address", &config.DatabaseConfig.DBAddr).Info("Database connected on address")

	routerService := router.NewRouterService(databaseService, log, &config.RouterConfig)
	return &App{
		dbService:     databaseService,
		routerService: routerService,
		log:           log,
	}

}

// Run the app
func (a *App) Run(addr string) {
	//TODO: add error handling - server crashing
	http.ListenAndServe(addr, a.routerService.Router)
}

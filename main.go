package main

import (
	"fmt"
	"net/http"

	"github.com/Sneh1999/Xpire/data"
	"github.com/Sneh1999/Xpire/models"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.Formatter = &logrus.JSONFormatter{}
	var config models.Config
	err := envconfig.Process("xpire", &config)

	if err != nil {
		log.WithError(err).Error("Couldn't load environment variables")
	}

	databaseService, err := data.NewDatabaseService(&config.DatabaseConfig, log)
	if err != nil {
		log.WithError(err).Fatal("Error setting up database service")
	}

	log.WithField("address", &config.DatabaseConfig.DBAddr).Info("Database connected on address")

	routerService := NewRouterService(databaseService, log, &config)

	log.Fatal(http.ListenAndServe(":"+fmt.Sprint(config.Port), routerService.Router))
}

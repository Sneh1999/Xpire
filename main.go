package main

import (
	"github.com/Sneh1999/Xpire/model"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.Formatter = &logrus.JSONFormatter{}
	log.Info("App Initialised")
	var config model.Config
	err := envconfig.Process("xpire", &config)
	
	if err != nil {
		log.WithError(err).Error("Couldn't load environment variables")
	}
	
	app := Initialize(&config.DatabaseConfig,log);

	app.Run(":8000")
}
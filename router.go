package main

import (
	"github.com/Sneh1999/Xpire/data"
	"github.com/Sneh1999/Xpire/models"
	"github.com/Sneh1999/Xpire/router"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type RouterService struct {
	Router       *mux.Router
	log          *logrus.Logger
	db           *data.DatabaseService
	routerConfig *models.RouterConfig
}

func NewRouterService(db *data.DatabaseService, log *logrus.Logger, routerConfig *models.RouterConfig) *RouterService {
	muxRouter := mux.NewRouter()
	authHandler := router.NewAuthHandler(db, log, &routerConfig.JWTConfig)
	muxRouter.HandleFunc("/v1/signup", authHandler.SignUp).Methods("POST")
	muxRouter.HandleFunc("/v1/login", authHandler.Login).Methods("POST")
	return &RouterService{
		Router:       muxRouter,
		log:          log,
		db:           db,
		routerConfig: routerConfig,
	}
}

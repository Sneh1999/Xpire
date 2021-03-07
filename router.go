package main

import (
	"github.com/Sneh1999/Xpire/data"
	"github.com/Sneh1999/Xpire/handlers"
	"github.com/Sneh1999/Xpire/middlewares"
	"github.com/Sneh1999/Xpire/models"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type RouterService struct {
	Router       *mux.Router
	log          *logrus.Logger
	db           *data.DatabaseService
	routerConfig *models.RouterConfig
}

func NewRouterService(db *data.DatabaseService, log *logrus.Logger, config *models.Config) *RouterService {
	muxRouter := mux.NewRouter()

	// Set up handlers
	authHandler := handlers.NewAuthHandler(db, log, &config.JWTConfig)
	orderHandler := handlers.NewOrderHandler(db, log, &config.JWTConfig)
	productHandler := handlers.NewProductHandler(db, log, &config.JWTConfig)
	// Set up middleware
	authMiddleware := middlewares.NewMiddlewareService(&config.JWTConfig, log)

	// Set up routes
	// Protected routes
	protectedRoutes := muxRouter.PathPrefix("/v1/api").Subrouter()
	protectedRoutes.Use(authMiddleware.AuthMiddleware)

	// Auth routes
	muxRouter.HandleFunc("/v1/signup", authHandler.SignUp).Methods("POST")
	muxRouter.HandleFunc("/v1/login", authHandler.Login).Methods("POST")

	// Product routes
	protectedRoutes.HandleFunc("/product", productHandler.GetProduct).Methods("GET")
	protectedRoutes.HandleFunc("/product", productHandler.CreateProduct).Methods("POST")
	protectedRoutes.HandleFunc("/product", productHandler.EditProduct).Methods("PUT")
	protectedRoutes.HandleFunc("/product", productHandler.DeleteProduct).Methods("DELETE")

	// Order routes
	// protectedRoutes.HandleFunc("/order", orderHandler.GetOrder).Methods("GET")
	protectedRoutes.HandleFunc("/order", orderHandler.CreateOrder).Methods("POST")
	// protectedRoutes.HandleFunc("/order", orderHandler.EditOrder).Methods("PUT")
	// protectedRoutes.HandleFunc("/order", orderHandler.DeleteOrder).Methods("DELETE")

	return &RouterService{
		Router:       muxRouter,
		log:          log,
		db:           db,
		routerConfig: &config.RouterConfig,
	}
}

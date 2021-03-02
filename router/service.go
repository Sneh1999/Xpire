package router

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type RouterService struct {
	Router *mux.Router
	log *logrus.Logger
}

func NewRouterService(log *logrus.Logger) *RouterService {
	router:= mux.NewRouter()

	// Authentication Routes
	// authenticationService := ""
	// // Product routes
	// productService := ""

	// router.HandleFunc("/", someFunction).Methods("GET")
	
	return &RouterService{
		Router: router,
		log: log,
	}
}


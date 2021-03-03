package router

import (
	"github.com/Sneh1999/Xpire/data"
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

func NewRouterService(db *data.DatabaseService, log *logrus.Logger, routerConfig *models.RouterConfig) *RouterService {
	router := mux.NewRouter()
	authHandler := NewAuthHandler(db, log, &routerConfig.JWTConfig)
	router.HandleFunc("/v1/signup", authHandler.SignUp).Methods("POST")
	return &RouterService{
		Router:       router,
		log:          log,
		db:           db,
		routerConfig: routerConfig,
	}
}

package router

import (
	"github.com/Sneh1999/Xpire/data"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type RouterService struct {
	Router *mux.Router
	log    *logrus.Logger
	db     *data.DatabaseService
}

func NewRouterService(log *logrus.Logger, db *data.DatabaseService) *RouterService {
	router := mux.NewRouter()
	return &RouterService{
		Router: router,
		log:    log,
		db:     db,
	}
}

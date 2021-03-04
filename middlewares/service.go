package middlewares

import (
	"github.com/Sneh1999/Xpire/models"
	"github.com/sirupsen/logrus"
)

type MiddlewareService struct {
	log       *logrus.Logger
	jwtConfig *models.JWTConfig
}

// NewMiddlewareService is used to initialize the middleware
func NewMiddlewareService(jwtConfig *models.JWTConfig, log *logrus.Logger) *MiddlewareService {

	middlwareService := &MiddlewareService{
		log:       log,
		jwtConfig: jwtConfig,
	}
	return middlwareService
}

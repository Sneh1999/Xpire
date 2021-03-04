package middlewares

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/Sneh1999/Xpire/models"
	"github.com/Sneh1999/Xpire/utils"
	"github.com/dgrijalva/jwt-go"
)

//AuthMiddleware
func (m *MiddlewareService) AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// log.Println("middleware", r.URL)
		j := &models.JwtWrapper{
			Issuer:          m.jwtConfig.Issuer,
			SecretKey:       m.jwtConfig.SecretKey,
			ExpirationHours: m.jwtConfig.ExpirationHours,
		}

		token := r.Header.Get("Authorization")

		if len(token) <= 7 {
			errResponse := &models.ErrorResponse{
				Message: "The user is unauthorized",
			}
			utils.WritePretty(w, http.StatusUnauthorized, errResponse)
			return
		}
		token = token[7:]
		claim, err := m.validateToken(token, j)
		if err != nil {
			errResponse := &models.ErrorResponse{
				Message: "The user is unauthorized",
			}
			utils.WritePretty(w, http.StatusUnauthorized, errResponse)
			return
		}

		ctx := context.WithValue(r.Context(), models.ContextKey("user_id"), claim.ID)

		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

//ValidateToken validates the token

func (m *MiddlewareService) validateToken(signedToken string, j *models.JwtWrapper) (*models.JwtClaim, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		// determines the structure for claims
		&models.JwtClaim{},
		// returns secret key
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.SecretKey), nil
		},
	)
	if err != nil {
		return nil, err
	}
	//converting token.claims to our claims model
	claims, ok := token.Claims.(*models.JwtClaim)

	if !ok {
		err = errors.New("Couldn't parse claims")
		return nil, err
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("JWT is expired")
		return nil, err
	}
	return claims, nil

}

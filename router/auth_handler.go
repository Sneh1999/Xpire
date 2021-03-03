package router

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/Sneh1999/Xpire/data"
	"github.com/Sneh1999/Xpire/models"
	"github.com/Sneh1999/Xpire/utils"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	db        *data.DatabaseService
	log       *logrus.Logger
	jwtConfig *models.JWTConfig
}

// NewAuthHandler helps in setting up the auth service
func NewAuthHandler(databaseService *data.DatabaseService, log *logrus.Logger, jwtConfig *models.JWTConfig) *AuthHandler {
	authHandler := &AuthHandler{
		db:        databaseService,
		log:       log,
		jwtConfig: jwtConfig,
	}
	return authHandler
}

// TODO: change the way router gets handled

// SignUp helps in adding a new user
func (auth *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	auth.log.WithFields(logrus.Fields{
		"email": user.Email,
		"name":  user.Name,
	}).Debug("Received Signup request")

	if err != nil {
		auth.log.WithError(err).Error("Invalid user details sent in request")
		utils.WritePretty(w, http.StatusBadRequest, "Send in the correct credentials")
		return
	}

	user.Password, err = auth.hashPassword(user.Password)

	if err != nil {
		auth.log.WithError(err).Error("Error in hashing the password")
		utils.WritePretty(w, http.StatusInternalServerError, "User couldnt be added")
		return
	}
	user.ID = uuid.NewV4()
	err = auth.db.AddUser(&user)
	if err != nil {
		auth.log.WithError(err).Error("Error in storing the user details")
		utils.WritePretty(w, http.StatusInternalServerError, "User couldnt be added")
		return
	}

	jwtWrapper := &models.JwtWrapper{
		SecretKey:       auth.jwtConfig.SecretKey,
		Issuer:          auth.jwtConfig.Issuer,
		ExpirationHours: auth.jwtConfig.ExpirationHours,
	}
	token, err := auth.generateToken(jwtWrapper, user.ID.String())

	if err != nil {
		auth.log.WithError(err).Error("Error in generating the jwt token")
		utils.WritePretty(w, http.StatusInternalServerError, "User couldnt be added")
		return
	}
	signUpResponse := &models.SignUpResponse{
		Token: token,
	}
	utils.WritePretty(w, http.StatusOK, signUpResponse)

}

//GenerateToken generates a token
func (auth *AuthHandler) generateToken(j *models.JwtWrapper, userID string) (string, error) {
	claims := &models.JwtClaim{
		ID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(j.ExpirationHours)).Unix(),
			Issuer:    j.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(j.SecretKey))

	if err != nil {
		return "", err
	}
	return signedToken, nil
}

//ValidateToken validates the token

func (auth *AuthHandler) validateToken(signedToken string, j *models.JwtWrapper) (*models.JwtClaim, error) {
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

// HashPassword hashes the password
func (auth *AuthHandler) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// CheckPassword decrypts the password
func (auth *AuthHandler) checkPassword(password string, userID string) error {
	id,err := uuid.FromString(userID)
	if err != nil {
		return err
	}
	user := &models.User{ID: id}
	err = auth.db.DB.Model(user).WherePK().Select()
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err
}

package router

import (
	"encoding/json"
	"net/http"

	"github.com/Sneh1999/Xpire/data"
	"github.com/Sneh1999/Xpire/model"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)


type AuthHandler struct {
	db  *data.DatabaseService
	log *logrus.Logger
}

// NewAuthHandler helps in setting up the auth service
func NewAuthHandler(databaseService *data.DatabaseService, log *logrus.Logger) *AuthHandler {
	authHandler := &AuthHandler{
		db:  databaseService,
		log: log,
	}
	return authHandler
}

// SignUp helps in adding a new user
func (auth *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)

	auth.log.WithFields(logrus.Fields{
		"email": user.Email,
		"name":  user.Name,
	}).Debug("Received Signup request")

	if err != nil {
		auth.log.WithError(err).Error("Invalid user details sent in request")
		WritePretty(w, http.StatusBadRequest, "Send in the correct credentials")
	}

	user.Password, err = auth.hashPassword(user.Password)

	if err != nil {
		auth.log.WithError(err).Error("Error in hashing the password")
		WritePretty(w, http.StatusInternalServerError, "User couldnt be added")
	}

	err = auth.db.AddUser(&user)
	if err != nil {
		auth.log.WithError(err).Error("Error in storing the user details")
		WritePretty(w, http.StatusInternalServerError, "User couldnt be added")
	}
	WritePretty(w, http.StatusOK, user)

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
func (auth *AuthHandler) checkPassword(password string, userID int64) error {
	user := &model.User{ID: userID}
	err := auth.db.DB.Model(user).WherePK().Select()
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err
}

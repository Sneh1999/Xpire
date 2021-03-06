package handlers

import (
	"fmt"
	"net/http"

	"github.com/Sneh1999/Xpire/data"
	"github.com/Sneh1999/Xpire/models"
	"github.com/Sneh1999/Xpire/utils"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type OrderHandler struct {
	db        *data.DatabaseService
	log       *logrus.Logger
	jwtConfig *models.JWTConfig
}

// NewOrderHandler helps in setting up the order service
func NewOrderHandler(databaseService *data.DatabaseService, log *logrus.Logger, jwtConfig *models.JWTConfig) *OrderHandler {
	orderHandler := &OrderHandler{
		db:        databaseService,
		log:       log,
		jwtConfig: jwtConfig,
	}
	return orderHandler
}

//CreateOrder helps in creating a new order
func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {

	var errorResponse models.ErrorResponse

	// Get the user which wants to  add the order
	userID := r.Context().Value(models.ContextKey("user_id"))
	fmt.Println(userID)
	user := &models.User{ID: userID.(string)}
	err := h.db.GetUser(user)

	if err != nil {
		h.log.WithError(err).Error("Invalid user details sent in request")
		errorResponse.Message = "Send in the correct credentials"
		utils.WritePretty(w, http.StatusBadRequest, &errorResponse)
		return
	}

	h.log.WithFields(logrus.Fields{
		"user": userID,
	}).Debug("Received Create Order request")

	// cerate the order
	order := &models.Order{
		UserID: userID.(string),
		ID:     uuid.NewV4().String(),
	}

	err = h.db.AddOrder(order)

	if err != nil {
		errorResponse.Message = "Error in creating the order"
		h.log.WithError(err).Error(errorResponse.Message)
		utils.WritePretty(w, http.StatusInternalServerError, &errorResponse)
		return
	}

	orderMessage := &models.OrderResponse{
		Message: "Order created succesfuly",
		OrderID: order.ID,
	}
	utils.WritePretty(w, http.StatusOK, orderMessage)
}

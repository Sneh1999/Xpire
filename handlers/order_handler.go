package handlers

import (
	"encoding/json"
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

//CreateOrder helps in retrieving an  order
func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	var errorResponse models.ErrorResponse
	var orderRequest models.GetOrderRequest

	err := json.NewDecoder(r.Body).Decode(&orderRequest)

	h.log.WithFields(logrus.Fields{
		"order": orderRequest,
	}).Info("Received Get Order request")

	order := &models.Order{
		ID: orderRequest.OrderId,
	}

	err = h.db.GetOrder(order)

	if err != nil {
		errorResponse.Message = "Given order doesnt exist"
		h.log.WithError(err).Error(errorResponse.Message)
		utils.WritePretty(w, http.StatusInternalServerError, &errorResponse)
		return
	}

	orderResponse := &models.GetOrderResponse{
		ID:       order.ID,
		Products: order.Products,
		UserID:   order.UserID,
	}
	utils.WritePretty(w, http.StatusOK, orderResponse)
}

//CreateOrder helps in creating a new order
func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {

	var errorResponse models.ErrorResponse

	// Get the user which wants to  add the order
	userID := r.Context().Value(models.ContextKey("user_id"))

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

	orderMessage := &models.CreateOrderResponse{
		Message: "Order created succesfuly",
		ID:      order.ID,
	}
	utils.WritePretty(w, http.StatusCreated, orderMessage)
}

//EsitOrder helps in editing an order
func (h *OrderHandler) EditOrder(w http.ResponseWriter, r *http.Request) {
	var errorResponse models.ErrorResponse
	var orderResponse models.EditOrderRequest

	err := json.NewDecoder(r.Body).Decode(&orderResponse)

	h.log.WithFields(logrus.Fields{
		"order": orderResponse,
	}).Info("Received Edit Order request")

	order := &models.Order{
		ID:       orderResponse.ID,
		Products: orderResponse.Products,
	}

	err = h.db.EditOrder(order)

	if err != nil {
		errorResponse.Message = "Given order couldnt be edited: " + err.Error()
		h.log.WithError(err).Error(errorResponse.Message)
		utils.WritePretty(w, http.StatusInternalServerError, &errorResponse)
		return
	}

	orderMessage := &models.EditOrderResponse{
		ID:       order.ID,
		Products: order.Products,
		UserID:   order.UserID,
	}
	utils.WritePretty(w, http.StatusOK, orderMessage)
}

//DeleteOrder helps in deleting an order
func (h *OrderHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	var errorResponse models.ErrorResponse
	var orderRequest models.DeleteOrderRequest
	err := json.NewDecoder(r.Body).Decode(&orderRequest)

	h.log.WithFields(logrus.Fields{
		"order": orderRequest,
	}).Info("Received Delete order request")

	order := &models.Order{
		ID: orderRequest.ID,
	}

	err = h.db.GetOrder(order)

	if err != nil {
		errorResponse.Message = "Given order doesnt exist"
		h.log.WithError(err).Error(errorResponse.Message)
		utils.WritePretty(w, http.StatusInternalServerError, &errorResponse)
		return
	}

	order = &models.Order{
		ID:       order.ID,
		UserID:   order.UserID,
		Products: order.Products,
		Delete:   true,
	}

	err = h.db.DeleteOrder(order)

	productMessage := &models.DeleteOrderResponse{
		Message: "Successfully deleted the order",
	}
	utils.WritePretty(w, http.StatusNoContent, productMessage)

}

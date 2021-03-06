package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Sneh1999/Xpire/data"
	"github.com/Sneh1999/Xpire/models"
	"github.com/Sneh1999/Xpire/utils"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type ProductHandler struct {
	db        *data.DatabaseService
	log       *logrus.Logger
	jwtConfig *models.JWTConfig
}

// NewProductHandler helps in setting up the order service
func NewProductHandler(databaseService *data.DatabaseService, log *logrus.Logger, jwtConfig *models.JWTConfig) *ProductHandler {
	productHandler := &ProductHandler{
		db:        databaseService,
		log:       log,
		jwtConfig: jwtConfig,
	}
	return productHandler
}

//CreateProduct helps in creating a new order
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var errorResponse models.ErrorResponse
	var productRequest models.ProductRequest
	// Get the order which wants to  add the order

	err := json.NewDecoder(r.Body).Decode(&productRequest)

	h.log.WithFields(logrus.Fields{
		"order": productRequest,
	}).Info("Received Create Order request")

	// cerate the order
	order := &models.Order{
		ID: productRequest.OrderID,
	}

	err = h.db.GetOrder(order)

	if err != nil {
		errorResponse.Message = "Given order doesnt exist"
		h.log.WithError(err).Error(errorResponse.Message)
		utils.WritePretty(w, http.StatusInternalServerError, &errorResponse)
		return
	}
	time, err := time.Parse("Mon Jan 2 15:04:05 -0700 MST 2006", productRequest.Expiry)
	if err != nil {
		errorResponse.Message = "Provided time is not correct"
		h.log.WithError(err).Error(errorResponse.Message)
		utils.WritePretty(w, http.StatusInternalServerError, &errorResponse)
		return
	}
	product := &models.Product{
		ID:      uuid.NewV4().String(),
		Name:    productRequest.Name,
		Expiry:  time,
		OrderID: productRequest.OrderID,
	}

	err = h.db.AddProduct(product)

	productMessage := &models.ProductResponse{
		Message:   "Product created succesfuly",
		ProductID: order.ID,
	}
	utils.WritePretty(w, http.StatusOK, productMessage)
}

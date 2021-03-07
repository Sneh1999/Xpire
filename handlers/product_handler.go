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

//GetProduct help in getting a product
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	var errorResponse models.ErrorResponse
	var productRequest models.GetProductRequest

	err := json.NewDecoder(r.Body).Decode(&productRequest)

	h.log.WithFields(logrus.Fields{
		"product": productRequest,
	}).Info("Received Get Product request")

	product := &models.Product{
		OrderID: productRequest.OrderID,
		ID:      productRequest.ProductID,
	}

	err = h.db.GetProduct(product)

	if err != nil {
		errorResponse.Message = "Given product doesnt exist"
		h.log.WithError(err).Error(errorResponse.Message)
		utils.WritePretty(w, http.StatusInternalServerError, &errorResponse)
		return
	}

	productMessage := &models.GetProductResponse{
		ID:      product.ID,
		Name:    product.Name,
		Expiry:  product.Expiry.String(),
		OrderID: product.OrderID,
	}
	utils.WritePretty(w, http.StatusOK, productMessage)

}

//CreateProduct helps in creating a new order
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var errorResponse models.ErrorResponse
	var productRequest models.CreateProductRequest
	// Get the order which wants to  add the order

	err := json.NewDecoder(r.Body).Decode(&productRequest)

	if err != nil {
		errorResponse.Message = "The details sent by the user are incorrect"
		h.log.WithError(err).Error(errorResponse.Message)
		utils.WritePretty(w, http.StatusInternalServerError, &errorResponse)
		return
	}

	h.log.WithFields(logrus.Fields{
		"order": productRequest,
	}).Info("Received Create Product request")

	// create the order
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

	productMessage := &models.CreateProductResponse{
		Message:   "Product created succesfuly",
		ProductID: product.ID,
	}
	utils.WritePretty(w, http.StatusOK, productMessage)
}

//EditProduct helps in edit a product
func (h *ProductHandler) EditProduct(w http.ResponseWriter, r *http.Request) {
	var errorResponse models.ErrorResponse
	var productRequest models.EditProductRequest

	err := json.NewDecoder(r.Body).Decode(&productRequest)

	h.log.WithFields(logrus.Fields{
		"product": productRequest,
	}).Info("Received Edit Product request")

	time, err := time.Parse("Mon Jan 2 15:04:05 -0700 MST 2006", productRequest.Expiry)

	if err != nil {
		errorResponse.Message = "Provided time is not correct"
		h.log.WithError(err).Error(errorResponse.Message)
		utils.WritePretty(w, http.StatusInternalServerError, &errorResponse)
		return
	}
	product := &models.Product{
		OrderID: productRequest.OrderID,
		Name:    productRequest.Name,
		Expiry:  time,
		ID:      productRequest.ProductID,
	}

	err = h.db.EditProduct(product)

	if err != nil {
		errorResponse.Message = "Given product couldnt be edited: " + err.Error()
		h.log.WithError(err).Error(errorResponse.Message)
		utils.WritePretty(w, http.StatusInternalServerError, &errorResponse)
		return
	}

	productMessage := &models.EditProductResponse{
		ID:      product.ID,
		Name:    product.Name,
		Expiry:  product.Expiry.String(),
		OrderID: product.OrderID,
	}
	utils.WritePretty(w, http.StatusOK, productMessage)
}

//DeleteProduct helps in deleting a product
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var errorResponse models.ErrorResponse
	var productRequest models.DeleteProductRequest

	err := json.NewDecoder(r.Body).Decode(&productRequest)

	h.log.WithFields(logrus.Fields{
		"product": productRequest,
	}).Info("Received Delete Product request")

	product := &models.Product{
		OrderID: productRequest.OrderID,
		ID:      productRequest.ProductID,
	}

	err = h.db.GetProduct(product)

	if err != nil {
		errorResponse.Message = "Given product doesnt exist"
		h.log.WithError(err).Error(errorResponse.Message)
		utils.WritePretty(w, http.StatusInternalServerError, &errorResponse)
		return
	}

	product = &models.Product{
		OrderID: productRequest.OrderID,
		ID:      productRequest.ProductID,
		Name:    product.Name,
		Expiry:  product.Expiry,
		Delete:  true,
	}

	err = h.db.DeleteProduct(product)

	if err != nil {
		errorResponse.Message = "Given product couldnt be deleted" + err.Error()
		h.log.WithError(err).Error(errorResponse.Message)
		utils.WritePretty(w, http.StatusInternalServerError, &errorResponse)
		return
	}

	productMessage := &models.DeleteProductResponse{
		Message: "Successfully deleted the product",
	}
	utils.WritePretty(w, http.StatusOK, productMessage)
}

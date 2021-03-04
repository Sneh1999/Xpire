package handlers

// import (
// 	"encoding/json"
// 	"net/http"

// 	"github.com/Sneh1999/Xpire/models"
// 	"github.com/Sneh1999/Xpire/utils"
// 	"github.com/sirupsen/logrus"
// )

// func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
// 	var orderRequest models.OrderRequest
// 	var errorResponse models.ErrorResponse
// 	err := json.NewDecoder(r.Body).Decode(&orderRequest)

// 	h.log.WithFields(logrus.Fields{
// 		"products": orderRequest.Products,
// 	}).Debug("Received Create Order request")

// 	if err != nil {
// 		h.log.WithError(err).Error("Invalid order details sent in request")
// 		errorResponse.Message = "Send in the correct credentials"
// 		utils.WritePretty(w, http.StatusBadRequest, &errorResponse)
// 		return
// 	}
// 	order := &models.Order{Products: orderRequest.Products,}
// }

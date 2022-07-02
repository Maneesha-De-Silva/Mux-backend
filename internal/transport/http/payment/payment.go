package payment

import (
	"backend/internal/types"
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

func (h *Handler) CreatePayment(w http.ResponseWriter, r *http.Request) {
	var newPayment types.Payment

	if err := json.NewDecoder(r.Body).Decode(&newPayment); err != nil {
		h.HandleErrorRespose(w, "Failed to decode JSON body", err, http.StatusInternalServerError)
		return
	}

	item, err := h.PaymentService.SavePayment(newPayment)

	if err != nil {
		h.HandleErrorRespose(w, "Unable to save the post", err, http.StatusBadRequest)
		return
	}

	h.HandleSuccessRespose(w, item)
}

// get all the payments
func (h *Handler) FetchPayments(w http.ResponseWriter, r *http.Request) {
	payments := h.PaymentService.GetPosts()
	h.HandleSuccessRespose(w, payments)
}

func (h *Handler) GetPaymentDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		h.HandleErrorRespose(w, "Unable to pass int", err, http.StatusBadRequest)
		return
	}

	details := h.PaymentService.GetPaymentByID(i)

	h.HandleSuccessRespose(w, details)

}

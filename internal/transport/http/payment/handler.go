package payment

import (
	"backend/internal/payment"
	"backend/internal/types"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type Handler struct {
	Router         *mux.Router
	PaymentService *payment.Service
}

func NewHandler(paymentService *payment.Service) *Handler {
	return &Handler{
		PaymentService: paymentService,
	}
}

func LogginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(
			log.Fields{
				"Method": r.Method,
				"Path":   r.URL.Path,
				"Host":   r.RemoteAddr,
			}).
			Info("Handeling Request")
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) SetupRotues() {
	h.Router = mux.NewRouter()

	h.Router.Use(LogginMiddleware)
	h.Router.HandleFunc("/payments/create", h.CreatePayment).Methods("POST")
	h.Router.HandleFunc("/payment/{id}", h.GetPaymentDetails).Methods("GET")
	h.Router.HandleFunc("/payment", h.FetchPayments).Methods("GET")

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		h.HandleSuccessRespose(w, struct {
			Message string `json:"message"`
		}{
			Message: "API Running okay",
		})
	})
}

// Handle Success Respose
func (h *Handler) HandleSuccessRespose(w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(resp)
}

// Handle Error Resposes
func (h *Handler) HandleErrorRespose(w http.ResponseWriter, message string, err error, errorCode int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(errorCode)

	if err := json.NewEncoder(w).Encode(types.ErrorResponse{
		Error:   message,
		Details: err.Error(),
	}); err != nil {
		panic(err)
	}
}

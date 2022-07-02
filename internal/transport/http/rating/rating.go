package rating

import (
//	"backend/internal/rating"
	"backend/internal/types"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *Handler) CreateRating(w http.ResponseWriter, r *http.Request) {
	var newRating types.Rating

	if err := json.NewDecoder(r.Body).Decode(&newRating); err != nil {
		h.HandleErrorRespose(w, "Failed to decode JSON body", err, http.StatusInternalServerError)
		return
	}

	rating, err := h.RatingService.SaveRating(newRating)

	if err != nil {
		h.HandleErrorRespose(w, "Unable to save the post", err, http.StatusBadRequest)
		return
	}

	h.HandleSuccessRespose(w, rating)
}

// get all the payments
func (h *Handler) FetchRating(w http.ResponseWriter, r *http.Request) {
	ratings := h.RatingService.GetPosts()
	h.HandleSuccessRespose(w, ratings)
}

func (h *Handler) GetRatingDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		h.HandleErrorRespose(w, "Unable to pass int", err, http.StatusBadRequest)
		return
	}

	details := h.RatingService.GetRatingByID(i)

	h.HandleSuccessRespose(w, details)

}

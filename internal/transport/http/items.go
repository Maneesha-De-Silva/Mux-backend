package http

import (
	"backend/internal/types"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *Handler) FetchItems(w http.ResponseWriter, r *http.Request) {
	items := h.ItemService.GetPosts()
	h.HandleSuccessRespose(w, items)
}

func (h *Handler) CreateItems(w http.ResponseWriter, r *http.Request) {
	var newItem types.Item

	if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
		h.HandleErrorRespose(w, "Failed to decode JSON body", err, http.StatusInternalServerError)
		return
	}

	item, err := h.ItemService.CreateItem(newItem)

	if err != nil {
		h.HandleErrorRespose(w, "Unable to save the post", err, http.StatusBadRequest)
		return
	}

	h.HandleSuccessRespose(w, item)
}

func (h *Handler) GetItemDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		h.HandleErrorRespose(w, "Unable to pass int", err, http.StatusBadRequest)
		return
	}

	details := h.ItemService.GetItemByID(i)

	h.HandleSuccessRespose(w, details)

}

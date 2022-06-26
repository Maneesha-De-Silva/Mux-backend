package http

import (
	"backend/internal/types"
	"encoding/json"
	"net/http"
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

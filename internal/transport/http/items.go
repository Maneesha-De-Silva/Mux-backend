package http

import "net/http"

func (h *Handler) FetchEvents(w http.ResponseWriter, r *http.Request) {
	items := h.ItemService.GetPosts()
	h.HandleSuccessRespose(w, items)
}

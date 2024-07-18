package journal

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matteoaricci/journal-be/types"
	"github.com/matteoaricci/journal-be/utils"
)

type Handler struct {
	store types.JournalStore
}

func NewHandler(store types.JournalStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/journals", h.getJournals).Methods("GET")
}

func (h *Handler) getJournals(w http.ResponseWriter, r *http.Request) {
	j, err := h.store.GetAllJournals()

	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	utils.WriteJSON(w, http.StatusAccepted, &j)
}
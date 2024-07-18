package journal

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matteoaricci/journal-be/types"
	"github.com/matteoaricci/journal-be/utils"
)

type Handler struct {
	store *types
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/journals", h.getJournal).Methods("GET")
}

func (h *Handler) getJournal(w http.ResponseWriter, r *http.Request) {
	var payload types.Journal
	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
}
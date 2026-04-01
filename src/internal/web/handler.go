package web

import (
	"encoding/json"
	"net/http"
	"tictac/internal/domain"
)

type Handler struct {
	processor domain.Processor
}

func NewHandler(p domain.Processor) *Handler {
	return &Handler{processor: p}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {

	var dto GameDTO
	json.NewDecoder(r.Body).Decode(&dto)
	updated, err := h.processor.Process(ToDomain(dto))
	if err != nil {
		json.NewEncoder(w).Encode(ToDTO(updated))
		http.Error(w, "invalid move", 400)

	} else {
		json.NewEncoder(w).Encode(ToDTO(updated))
	}

}

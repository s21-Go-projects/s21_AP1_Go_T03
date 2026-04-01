package web

import (
	"encoding/json"
	"net/http"
	"strings"

	"tictac/internal/domain"

	"github.com/google/uuid"
)

type Handler struct {
	service domain.Service
}

func NewHandler(s domain.Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) Move(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/game/")
	id, _ := uuid.Parse(idStr)

	var dto GameDTO
	json.NewDecoder(r.Body).Decode(&dto)

	prev, _ := h.service.Get(id)
	next := ToDomain(dto)
	next.ID = id

	if !h.service.Validate(prev, next) {
		http.Error(w, "invalid move", 400)
		return
	}

	h.service.Save(next)

	if h.service.CheckWinner(next.Board) != domain.Empty {
		json.NewEncoder(w).Encode(ToDTO(next))
		return
	}

	updated, _ := h.service.NextMove(next)
	h.service.Save(updated)

	json.NewEncoder(w).Encode(ToDTO(updated))
}

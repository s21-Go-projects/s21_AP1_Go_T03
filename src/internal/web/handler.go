package web

import (
	"encoding/json"
	"net/http"
	"strings"

	"webServerKrest/internal/domain"
)

type GameHandler struct {
	service domain.GameService
}

func NewGameHandler(s domain.GameService) *GameHandler {
	return &GameHandler{service: s}
}

func (h *GameHandler) HandleGame(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/game/")

	var req GameRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	game := ToDomain(id, req)

	if err := h.service.Validate(game); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	game = h.service.NextMove(game)

	resp := ToResponse(game)
	json.NewEncoder(w).Encode(resp)
}

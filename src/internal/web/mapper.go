package web

import (
	"tictac/internal/domain"

	"github.com/google/uuid"
)

func ToDTO(g domain.Game) GameDTO {
	return GameDTO{ID: g.ID.String(), Board: g.Board, Winner: g.Winner}
}

func ToDomain(dto GameDTO) domain.Game {
	id, _ := uuid.Parse(dto.ID)
	return domain.Game{ID: id, Board: dto.Board, Winner: dto.Winner}
}

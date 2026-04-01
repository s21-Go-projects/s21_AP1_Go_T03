package web

import "tictac/internal/domain"

func ToDTO(g domain.Game) GameDTO {
	return GameDTO{ID: g.ID.String(), Board: g.Board}
}

func ToDomain(dto GameDTO) domain.Game {
	return domain.Game{Board: dto.Board}
}

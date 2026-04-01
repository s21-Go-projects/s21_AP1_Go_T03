package datasource

import "tictac/internal/domain"

func toEntity(g domain.Game) GameData {
	return GameData{ID: g.ID, Board: g.Board, Winner: g.Winner}
}

func toDomain(gd GameData) domain.Game {
	return domain.Game{ID: gd.ID, Board: gd.Board, Winner: gd.Winner}
}

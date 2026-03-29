package datasource

import "webServerKrest/internal/domain"

func ToDTO(g domain.Game) GameDTO {
	board := make([][]int, len(g.Board))
	for i := range g.Board {
		board[i] = make([]int, len(g.Board[i]))
		for j := range g.Board[i] {
			board[i][j] = int(g.Board[i][j])
		}
	}
	return GameDTO{ID: g.ID, Board: board}
}

func ToDomain(g GameDTO) domain.Game {
	board := make([][]domain.Cell, len(g.Board))
	for i := range g.Board {
		board[i] = make([]domain.Cell, len(g.Board[i]))
		for j := range g.Board[i] {
			board[i][j] = domain.Cell(g.Board[i][j])
		}
	}
	return domain.Game{ID: g.ID, Board: board}
}

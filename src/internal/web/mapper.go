package web

import "webServerKrest/internal/domain"

func ToDomain(id string, req GameRequest) domain.Game {
	board := make([][]domain.Cell, len(req.Board))
	for i := range req.Board {
		board[i] = make([]domain.Cell, len(req.Board[i]))
		for j := range req.Board[i] {
			board[i][j] = domain.Cell(req.Board[i][j])
		}
	}
	return domain.Game{ID: id, Board: board}
}

func ToResponse(g domain.Game) GameResponse {
	board := make([][]int, len(g.Board))
	for i := range g.Board {
		board[i] = make([]int, len(g.Board[i]))
		for j := range g.Board[i] {
			board[i][j] = int(g.Board[i][j])
		}
	}
	return GameResponse{ID: g.ID, Board: board}
}

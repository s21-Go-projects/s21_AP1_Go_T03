package datasource

import "github.com/google/uuid"

type GameData struct {
	ID     uuid.UUID
	Board  [3][3]int
	Winner int
}

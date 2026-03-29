package di

import (
	"webServerKrest/internal/domain"
	//"webServerKrest/repository"
	"webServerKrest/internal/web"

	"webServerKrest/internal/datasource"
)

func NewStorage() *datasource.Storage {
	return datasource.NewStorage()
}

// func NewRepo(s *datasource.Storage) repository.GameRepository {
// 	return repository.NewGameRepository(s)
// }

func NewService() domain.GameService {
	return domain.NewGameService()
}

func NewHandler(s domain.GameService) *web.GameHandler {
	return web.NewGameHandler(s)
}

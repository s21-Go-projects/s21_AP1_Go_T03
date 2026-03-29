package di

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		NewStorage,
		//	NewRepo,
		NewService,
		NewHandler,
	),
)

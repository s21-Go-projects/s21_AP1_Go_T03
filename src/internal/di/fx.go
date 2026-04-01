package di

import (
	"context"
	"net/http"
	"tictac/internal/datasource"
	"tictac/internal/domain"
	"tictac/internal/web"

	"go.uber.org/fx"
)

func Run(handler *web.Handler) {
	mux := http.NewServeMux()
	mux.HandleFunc("/game/", handler.Move)
	http.ListenAndServe(":8080", mux)
}

var Module = fx.Options(
	fx.Provide(
		datasource.NewStorage,
		datasource.NewRepository,
		func(r datasource.Repository) domain.Repository { return r },
		domain.NewService,
		web.NewHandler,
	),
	fx.Invoke(func(lc fx.Lifecycle, handler *web.Handler) {
		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				go Run(handler)
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return nil
			},
		})
	}),
)

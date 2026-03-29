package main

import (
	"net/http"
	"webServerKrest/internal/di"
	"webServerKrest/internal/web"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		di.Module,
		fx.Invoke(func(handler *web.GameHandler) {
			http.HandleFunc("/game/", handler.HandleGame)
			go http.ListenAndServe(":8080", nil)
		}),
	).Run()
}

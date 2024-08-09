package main

import (
	"net/http"
	"voting_system/handler"
	"voting_system/internal"

	"github.com/go-chi/chi"
)

func main() {
	handler := handler.NewUserClient()
	r := chi.NewRouter()
	r.Post("/token", internal.GenerateToken)
	r.Post("/user/register", handler.SetUser)
	r.Group(func(r chi.Router) {
		r.Post("/user/castvote", internal.ValidateTokenMiddleware(handler.CastVote))
	})

	http.ListenAndServe(":3009", r)
}

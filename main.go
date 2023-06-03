package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"westminster/application"
	"westminster/application/adapter/rest"
	"westminster/application/adapter/rest/handler"
)

func main() {
	_ = application.New()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	memberHandler := handler.NewMemberHandler(nil)

	rest.RegisterRoutes(r, memberHandler)

	r.Post("/members", memberHandler.CreateMember)
	http.ListenAndServe(":3000", r)
}

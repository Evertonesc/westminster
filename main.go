package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"westminster/application"
	"westminster/application/adapter/rest/handler"
)

func main() {
	app := application.New()
	fmt.Println(app)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome from chi rounter!"))
	})

	memberHandler := handler.NewMemberHandler(nil)

	r.Post("/members", memberHandler.CreateMember)
	http.ListenAndServe(":3000", r)
}

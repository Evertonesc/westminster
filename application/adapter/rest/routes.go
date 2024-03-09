package rest

import (
	"github.com/go-chi/chi/v5"

	"westminster/application/adapter/rest/handler"
)

func RegisterRoutes(r chi.Router, memberHandler *handler.MemberHandler) {
	r.Post("/members", memberHandler.CreateMember)
}

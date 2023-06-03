package handler

//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE

import (
	"context"
	"net/http"

	"github.com/go-chi/render"

	"westminster/application/domain"
	"westminster/application/usecase"
)

type CreateMemberUseCase interface {
	Execute(ctx context.Context, i usecase.MemberInteractor) (domain.Member, error)
}

type MemberHandler struct {
	uc CreateMemberUseCase
}

type MemberResponse struct {
	Name string `json:"name"`
}

func NewMemberHandler(uc CreateMemberUseCase) MemberHandler {
	return MemberHandler{
		uc: uc,
	}
}

func (h MemberHandler) CreateMember(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, MemberResponse{Name: "some name"})
}

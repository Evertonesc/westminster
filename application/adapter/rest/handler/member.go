package handler

import (
	"context"
	"net/http"

	"github.com/go-chi/render"

	"westminster/application/adapter/rest/presenter"
	"westminster/application/domain"
	"westminster/application/usecase"
)

type CreateMemberUseCase interface {
	Execute(ctx context.Context, i usecase.MemberInteractor) (domain.Member, error)
}

type MemberHandler struct {
	uc CreateMemberUseCase
}

func NewMemberHandler(uc CreateMemberUseCase) *MemberHandler {
	return &MemberHandler{
		uc: uc,
	}
}

func (h *MemberHandler) CreateMember(w http.ResponseWriter, r *http.Request) {
	var req presenter.MemberRequest
	render.Bind(r, &req)

	render.Status(r, http.StatusCreated)
	render.Render(w, r, &presenter.MemberResponse{
		Name: req.Name,
	})
}

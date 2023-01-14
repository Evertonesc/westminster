package usecase

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"westminster/application/adapter/rest/presenter"
	"westminster/application/domain"
)

func TestCreateMemberUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	now := time.Now().UTC()

	type fields struct {
		w func() domain.MemberWriter
	}
	type args struct {
		ctx context.Context
		mr  presenter.MemberRequest
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantErr    error
		wantMember domain.Member
	}{
		{
			name: "should successfully create a member",
			fields: fields{
				w: func() domain.MemberWriter {
					mock := domain.NewMockMemberWriter(ctrl)
					mock.EXPECT().
						CreateMember(gomock.Any(), domain.Member{
							Name:            "some name",
							FinancialNumber: 434,
							Location:        "some location",
							Enabled:         true,
						}).
						Return(domain.Member{
							ID:              "some-id",
							Name:            "some name",
							FinancialNumber: 434,
							Location:        "some location",
							Enabled:         true,
							CreatedAt:       now,
							UpdatedAt:       now,
						}, nil)
					return mock
				},
			},
			args: args{
				ctx: context.Background(),
				mr: presenter.MemberRequest{
					Name:            "some name",
					Location:        "some location",
					FinancialNumber: 434,
					Enabled:         true,
				},
			},
			wantMember: domain.Member{
				ID:              "some-id",
				Name:            "some name",
				FinancialNumber: 434,
				Location:        "some location",
				Enabled:         true,
				CreatedAt:       now,
				UpdatedAt:       now,
			},
			wantErr: nil,
		},
		{
			name: "should return an error when the database writer fails",
			fields: fields{
				w: func() domain.MemberWriter {
					mock := domain.NewMockMemberWriter(ctrl)
					mock.EXPECT().
						CreateMember(gomock.Any(), domain.Member{
							Name:            "some name",
							FinancialNumber: 434,
							Location:        "some location",
							Enabled:         true,
						}).
						Return(domain.Member{}, errors.New("unexpected error"))
					return mock
				},
			},
			args: args{
				ctx: context.Background(),
				mr: presenter.MemberRequest{
					Name:            "some name",
					Location:        "some location",
					FinancialNumber: 434,
					Enabled:         true,
				},
			},
			wantMember: domain.Member{},
			wantErr:    errors.New("unexpected error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			uc := NewCreateMemberUseCase(tt.fields.w())
			m, err := uc.Execute(tt.args.ctx, tt.args.mr)

			assert.EqualValues(t, tt.wantErr, err)
			assert.EqualValues(t, tt.wantMember, m)
		})
	}
}

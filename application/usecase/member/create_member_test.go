package usecase

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"orinz/application/adapter/rest/presenter"
	"orinz/application/domain/member"
	"testing"
)

func TestCreateMemberUseCase_Execute(t *testing.T) {
	type fields struct {
		repository func() member.Repository
	}
	type args struct {
		ctx           context.Context
		memberRequest presenter.MemberRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "should return error when repository fails to insert a member",
			fields: fields{
				repository: func() member.Repository {
					mock := member.NewMockRepository(gomock.NewController(t))
					mock.EXPECT().
						CreateMember(gomock.Any(), member.Member{
							Name:      "Clive Staples Lewis",
							BirthDate: "1989-11-29",
							Address:   "Heaven street",
							Email:     "cslewis@narnia.com",
						}).
						Return(errors.New("error when trying to create a member"))
					return mock
				},
			},
			args: args{
				ctx: context.Background(),
				memberRequest: presenter.MemberRequest{
					Name:      "Clive Staples Lewis",
					BirthDate: "1989-11-29",
					Address:   "Heaven street",
					Email:     "cslewis@narnia.com",
				},
			},
			wantErr: errors.New("error when trying to create a member"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := NewCreateMemberUseCase(tt.fields.repository())
			err := uc.Execute(tt.args.ctx, tt.args.memberRequest)
			if err != nil {
				assert.Equalf(t, tt.wantErr, err, "Execute() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

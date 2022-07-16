package usecase

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"orinz/application/domain/class"
	"testing"
)

func TestCreateClassUsecase_CreateClass(t *testing.T) {
	type fields struct {
		r func() class.Repository
	}
	type args struct {
		ctx       context.Context
		className class.Class
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "should fail when repository returns an error",
			fields: fields{
				r: func() class.Repository {
					mock := class.NewMockRepository(gomock.NewController(t))
					mock.EXPECT().
						Create(gomock.Any(), class.Class{Name: "some-class-name"}).
						Return(errors.New("unexpected error"))
					return mock
				},
			},
			args: args{
				ctx: context.Background(),
				className: class.Class{
					Name: "some-class-name",
				},
			},
			wantErr: errors.New("unexpected error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := New(tt.fields.r())
			if err := uc.Execute(tt.args.ctx, tt.args.className); err != nil {
				assert.Equalf(t, tt.wantErr, err, "Execute() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

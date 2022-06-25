package usecase

import (
	"context"
	"orinz/application/adapter/rest/presenter"
	"orinz/application/domain/member"
	"testing"
)

func TestCreateMemberUseCase_Execute(t *testing.T) {
	type fields struct {
		repository member.Repository
	}
	type args struct {
		ctx           context.Context
		memberRequest presenter.MemberRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := CreateMemberUseCase{
				repository: tt.fields.repository,
			}
			if err := uc.Execute(tt.args.ctx, tt.args.memberRequest); (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

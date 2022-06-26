package repository

import (
	"context"
	"orinz/application/domain/member"
)

type MemberRepository struct {
}

func NewMemberRepository() MemberRepository {
	return MemberRepository{}
}

func (r MemberRepository) CreateMember(ctx context.Context, member member.Member) error {
	return nil
}

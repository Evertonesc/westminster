package repository

import (
	"context"
	"errors"
	"fmt"
	"orinz/application/domain/member"
)

type Member struct {
	Name      string
	BirthDate string
	Address   string
	Email     string
}

func NewRepositoryMember(member member.Member) Member {
	return Member{
		Name:      member.Name,
		BirthDate: member.BirthDate,
		Address:   member.Address,
		Email:     member.Email,
	}
}

type MemberRepository struct {
}

func NewMemberRepository() MemberRepository {
	return MemberRepository{}
}

func (r MemberRepository) CreateMember(ctx context.Context, member member.Member) error {
	m := NewRepositoryMember(member)
	fmt.Println(m)

	return errors.New("error when trying to create a member")
}

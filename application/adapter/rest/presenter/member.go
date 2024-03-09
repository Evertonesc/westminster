package presenter

import (
	"errors"
	"net/http"
)

type MemberRequest struct {
	Name            string `json:"name"`
	Location        string `json:"location"`
	FinancialNumber int    `json:"financial_number"`
	Enabled         bool   `json:"enabled"`
}

type MemberResponse struct {
	Name string `json:"name"`
}

func (mr MemberRequest) Bind(r *http.Request) error {
	return nil
}

func (mr MemberRequest) Validate() error {
	if mr.Name == "" {
		return errors.New("name cannot be empty")
	}

	if mr.Location == "" {
		return errors.New("location cannot be empty")
	}

	if mr.FinancialNumber <= 0 {
		return errors.New("financial number must be a non zero value")
	}

	if !mr.Enabled {
		return errors.New("cannot create a disable member")
	}

	return nil
}

func (mrs *MemberResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

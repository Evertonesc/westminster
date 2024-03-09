package presenter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemberRequest_Validate(t *testing.T) {
	type fields struct {
		Name            string
		Location        string
		FinancialNumber int
		Enabled         bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "should return a validation error when the member name is empty",
			fields: fields{
				Name: "",
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Error(t, err, "name cannot be empty")
			},
		},
		{
			name: "should return a validation error when the location is empty",
			fields: fields{
				Name:     "Some name",
				Location: "",
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Error(t, err, "location cannot be empty")
			},
		},
		{
			name: "should return a validation error when the financial number is zero",
			fields: fields{
				Name:            "Some name",
				Location:        "some location",
				FinancialNumber: 0,
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Error(t, err, "financial number must be a non zero value")
			},
		},
		{
			name: "should return a validation error when enabled property is false",
			fields: fields{
				Name:            "Some name",
				Location:        "some location",
				FinancialNumber: 0,
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Error(t, err, "enabled cannot be false")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mr := MemberRequest{
				Name:            tt.fields.Name,
				Location:        tt.fields.Location,
				FinancialNumber: tt.fields.FinancialNumber,
				Enabled:         tt.fields.Enabled,
			}
			err := mr.Validate()
			if !tt.wantErr(t, err) {
				t.Errorf("MemberRequest.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

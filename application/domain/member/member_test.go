package member

import (
	"testing"
)

func TestMember_Validate(t *testing.T) {
	type fields struct {
		Name      string
		BirthDate string
		Address   string
		Email     string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "should return invalid member when name is empty",
			fields: fields{
				Name:      "",
				BirthDate: "1834-06-19",
				Address:   "England",
				Email:     "charles@baptist.com",
			},
			want: false,
		},
		{
			name: "should return invalid member when address is empty",
			fields: fields{
				Name:      "Charles Spurgeon",
				BirthDate: "1834-06-19",
				Address:   "",
				Email:     "charles@baptist.com",
			},
			want: false,
		},

		{
			name: "should return invalid member when birthdate is empty",
			fields: fields{
				Name:      "Charles Spurgeon",
				BirthDate: "",
				Address:   "England",
				Email:     "charles@baptist.com",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := New(tt.fields.Name, tt.fields.BirthDate, tt.fields.Address, tt.fields.Email)
			if got := m.Validate(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

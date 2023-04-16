package presenter

type MemberRequest struct {
	Name            string `json:"name"`
	Location        string `json:"location"`
	FinancialNumber int    `json:"financial_number"`
	Enabled         bool   `json:"enabled"`
}

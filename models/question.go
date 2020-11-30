package models

// Question struct used for benefit dependency information
type Question struct {
	ID           string   `json:"id"`
	Answer       string   `json:"answer"`
	Description  string   `json:"description"`
	OpenFiscaIds []string `json:"openfiscaids"`
}

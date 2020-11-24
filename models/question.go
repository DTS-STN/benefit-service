package models

// Question struct used for benefit dependency information
type Question struct {
	ID           string   `json:"id"`
	Answer       string   `json:"answer"`
	Description  string   `json:"description"`
	// List of which benefits the question is a dependency of
	DependencyOf []string `json:"dependency_of"`
	// The OpenFisca input variable name(s) attributed to this question
	OpenFiscaIds []string `json:"openfisca_ids"`
}

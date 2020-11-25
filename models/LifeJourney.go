package models

type LifeJourney struct {
	ID              int      `json:"id"`
	Title           string   `json:"title"`
	Description     string   `json:"description"`
	Title_fr        string   `json:"title_fr"`
	Description_fr  string   `json:"description_fr"`
	RelatedBenefits []string `json:"related_benefits"`
}

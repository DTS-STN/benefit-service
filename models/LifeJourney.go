package models

type LifeJourney struct {
	ID              string         `json:"id"`
	Title           string         `json:"title"`
	Description     string         `json:"description"`
	RelatedBenefits []string       `json:"related_benefits"`
	BenefitDetails  []FieldDetails `json:"lifejourney_details"`
}

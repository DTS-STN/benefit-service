package models

type Benefits struct {
	ID              string         `json:"id"`
	Title           string         `json:"title"`
	Description     string         `json:"description"`
	RelatedBenefits []string       `json:"related_benefits"`
	BenefitDetails  []FieldDetails `json:"benefit_details"`
}

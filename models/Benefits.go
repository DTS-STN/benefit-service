package models

type Benefits struct {
	ID              string         `json:"id"`
	Title           string         `json:"title"`
	TitleFr         string         `json:"title_fr"`
	Description     string         `json:"description"`
	DescriptionFr   string         `json:"description_fr"`
	RelatedBenefits []string       `json:"related_benefits"`
	BenefitDetails  []FieldDetails `json:"benefit_details"`
}

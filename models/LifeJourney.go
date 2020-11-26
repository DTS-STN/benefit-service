package models

type LifeJourney struct {
	ID              int            `json:"id"`
	Title           string         `json:"title"`
	Description     string         `json:"description"`
	TitleFr         string         `json:"title_fr"`
	DescriptionFr   string         `json:"description_fr"`
	RelatedBenefits []string       `json:"related_benefits"`
	BenefitDetails  []FieldDetails `json:"lifejourney_details"`
}

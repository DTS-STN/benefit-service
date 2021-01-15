package models

type Benefits struct {
	ID              int      `json:"id,string"`
	Title           string   `json:"title"`
	Description     string   `json:"description"`
	LongDescription string   `json:"long_description"`
	RelatedBenefits []string `json:"related_benefits"`
	ServiceType     string   `json:"service_type"`
	BenefitType     string   `json:"benefit_type"`
	BenefitKey      string   `json:"benefit_key"`
	BenefitTags     []string `json:"benefit_tags"`
	RedirectURL     string   `json:"redirect_url"`
	APIURL          string   `json:"api_url"`
}

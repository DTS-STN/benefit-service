package benefits

import (
	"github.com/DTS-STN/benefit-service/models"
)

type BenefitsInterface interface {
	Benefits(lang string) []models.Benefits
	Benefit(lang, id string) (models.Benefits, error)
	LoadBenefits(lang string) (Benefits []models.Benefits, err error)
}

type BenefitsServiceStruct struct {
}

var BenefitsService BenefitsInterface = new(BenefitsServiceStruct)

package benefits

import (
	"github.com/DTS-STN/benefit-service/models"
)

type BenefitsInterface interface {
	Benefits(lang string) []models.Benefits
	LoadBenefits(lang string) ([]models.Benefits, error)
	GetBenefitById(lang, benefitId string) (models.Benefits, error)
}

type ServiceStruct struct {
}

var Service BenefitsInterface = new(ServiceStruct)

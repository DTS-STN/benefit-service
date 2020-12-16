package benefits

import (
	"github.com/DTS-STN/benefit-service/models"
)

type BenefitsInterface interface {
	GetAll(lang string) []models.Benefits
	LoadBenefits(lang string) ([]models.Benefits, error)
	GetByID(lang, benefitId string) (models.Benefits, error)
}

type ServiceStruct struct {
}

var Service BenefitsInterface = new(ServiceStruct)

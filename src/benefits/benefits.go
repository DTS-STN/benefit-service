package benefits

import (
	"github.com/DTS-STN/benefit-service/models"
)

type BenefitsInterface interface {
	Benefits() []models.Benefits
	LoadBenefits() (Benefits []models.Benefits, err error)
}

type BenefitsServiceStruct struct {
	Filename string
}

var BenefitsService BenefitsInterface

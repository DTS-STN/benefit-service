package benefits

import (
	"github.com/DTS-STN/benefit-service/models"
)

type BenefitsInterface interface {
	Benefits() []models.Benefits
	Benefit(id string) (models.Benefits, error)
	LoadBenefits() (Benefits []models.Benefits, err error)
	ClearBenefits()
	SetFilePath(path string)
}

type BenefitsServiceStruct struct {
	Filename string
}

var BenefitsService BenefitsInterface = new(BenefitsServiceStruct)

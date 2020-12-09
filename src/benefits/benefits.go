package benefits

import (
	"github.com/DTS-STN/benefit-service/models"
)

type BenefitsInterface interface {
	Benefits() []models.Benefits
	LoadBenefits() (Benefits []models.Benefits, err error)
	SetFilePath(path string)
}

type BenefitsServiceStruct struct {
	Filename string
}

var BenefitsService BenefitsInterface = new(BenefitsServiceStruct)

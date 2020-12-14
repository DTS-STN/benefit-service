package benefits

import (
	"github.com/DTS-STN/benefit-service/models"
)

type BenefitsInterface interface {
	Benefits() []models.Benefits
	LoadBenefits() (Benefits []models.Benefits, err error)
	GetBenefitById(benefitId string) (benefit models.Benefits, err error)
}

type BenefitsServiceStruct struct {
	Filename string
}

var BenefitsService BenefitsInterface

func (q BenefitsServiceStruct) GetBenefitById(benefitId string) (benefit models.Benefits, err error) {

	benList, err := BenefitsService.LoadBenefits()
	if err != nil {
		return benefit, err
	}

	for _, ben := range benList {
		if ben.ID == benefitId {
			benefit = ben
		}
	}
	return benefit, nil
}

package benefits

import (
	"github.com/DTS-STN/benefit-service/models"
)

type BenefitsInterface interface {
	GetAll() ([]models.Benefits, error)
	GetByID(id string) (models.Benefits, error)
}

type ServiceStruct struct {
	repository BenefitsInterface
}

// GetAll returns all Benefits
func (q *ServiceStruct) GetAll() ([]models.Benefits, error) {
	return q.repository.GetAll()
}

// GetByID returns a Benefit from an ID
func (q *ServiceStruct) GetByID(id string) (models.Benefits, error) {
	return q.repository.GetByID(id)
}

func NewService(repository BenefitRepo) *ServiceStruct {
	return &ServiceStruct{repository: repository}
}

var Service BenefitsInterface

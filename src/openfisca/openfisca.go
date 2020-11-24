package openfisca

import (
	"github.com/DTS-STN/benefit-service/bindings"
	"github.com/DTS-STN/benefit-service/renderings"

	"bytes"
	"encoding/json"
	"net/http"
)

// OFInterface mainly used for testing
type OFInterface interface {
	SendRequest(request *bindings.BenefitQuestionsRequest) (renderings.BenefitQuestionsResponse, error)
}

// OFService ...
// Generic struct
type OFService struct{}

// Service that others can use to interact with OpenFisca functions
var Service OFInterface

// SendRequest ...
// OpenFisca interface method for sending requests to OpenFisca
func (of OFService) SendRequest(QuestionsRequest *bindings.BenefitQuestionsRequest) (renderings.BenefitQuestionsResponse, error) {

	//Modify QuestionsRequest if necessary
	requestBody, err := json.Marshal(QuestionsRequest)
	if err != nil {
		return renderings.BenefitQuestionsResponse{}, err
	}

	//TODO: Put url in a config
	resp, err := http.Post("https://fd7a43f1-b30f-4895-836d-5b52cede5318.mock.pstmn.io/trace", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return renderings.BenefitQuestionsResponse{}, err
	}

	// Defer the closing of the body
	defer resp.Body.Close()
	temp := &renderings.BenefitQuestionsResponse{}
	err = json.NewDecoder(resp.Body).Decode(temp)

	return *temp, err
}

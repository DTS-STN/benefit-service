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
func (of OFService) SendRequest(OpenFiscaRequest map[string]interface{}) (OpenFiscaResponse map[string]interface{}, err error) {

	//Modify QuestionsRequest if necessary
	requestBody, err := json.Marshal(OpenFiscaRequest)
	if err != nil {
		return OpenFiscaResponse, err
	}

	//TODO: Put url in a config
	resp, err := http.Post("http://localhost:5000/trace", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return OpenFiscaResponse, err
	}

	// Defer the closing of the body
	defer resp.Body.Close()
	temp := &OpenFiscaResponse
	err = json.NewDecoder(resp.Body).Decode(temp)

	return *temp, err
}

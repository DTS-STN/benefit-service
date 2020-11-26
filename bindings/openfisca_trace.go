package bindings

// OpenFiscaTraceRequest is the request sent by the Benefit Service to OpenFisca that contains
//  the benefits for which to find the list of dependencies.
type OpenFiscaTraceRequest struct {
	// Date period for request in ms since epoch
	RequestDate int64 `json:"request_date"`
	// Array of life journeys, which represent a subset of benefits.
	// Questions for the benefits under the life journey will be returned.
	LifeJourneys []string `json:"life_journeys"`
	// Array of specific benefits for which to get the questions.
	BenefitList []string `json:"benefit_list"`
}

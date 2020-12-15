package bindings

// BenefitQuestionsRequest is the request sent by the client that contains the benefits
// for which they would like the list of questions/dependencies.
// swagger:model BenefitQuestionsRequest
type BenefitQuestionsRequest struct {
	// Date period for request in ms since epoch
	RequestDate int64 `json:"request_date" example:"1608054805"`
	// Array of life journeys, which represent a subset of benefits.
	// Questions for the benefits under the life journey will be returned.
	LifeJourneys []string `json:"life_journeys"`
	// Array of specific benefits for which to get the questions.
	BenefitList []string `json:"benefit_list" example:"seniors_card__is_eligible, student_card__is_eligible"`
}

package bindings

type BenefitEligibilityRequest struct {
	IncomeDetails      string `json:"incomeDetails"`
	TimeOutOfWork      string `json:"timeOutOfWork"`
	AbleToWork         string `json:"ableToWork"`
	ReasonForOutOfWork string `json:"reasonForOutOfWork"`
	Gender             string `json:"gender"`
	Lang			   string `json:"lang"`
}

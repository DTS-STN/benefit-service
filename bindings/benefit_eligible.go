package bindings

type BenefitEligibilityRequest struct {
	IncomeDetails       string `json:"incomeDetails"`
	OutOfWork           string `json:"outOfWork"`
	AbleToWork          string `json:"ableToWork"`
	ReasonForSeparation string `json:"reasonForSeparation"`
	Gender              string `json:"gender"`
	Province            string `json:"province"`
}

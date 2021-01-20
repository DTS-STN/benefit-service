package bindings

type BenefitApplyRequest struct {
	BenefitType           string `json:"benefitType"`
	IncomeDetails         string `json:"incomeDetails"`
	OutOfWork             string `json:"outOfWork"`
	ReasonForSeperation   string `json:"reasonForSeperation"`
	RegularLookingForWork string `json:"regularLookingForWork"`
}

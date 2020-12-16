package bindings

type (
	// BenefitRequest is the request sent by the client that contains the information
	// required for the Benefit Service to return details Benefits.
	BenefitsRequest struct {
		Id     string `json:"id"`
		IdList string `json:"id_list"`
		Lang   string `json:"lang"`
	}
)

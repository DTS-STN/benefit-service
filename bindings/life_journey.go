package bindings

type (
	// LifeJourneyRequest is the request sent by the client that contains the information
	// required for the Benefit Service to return details on a Life Journey.
	LifeJourneyRequest struct {
		// Date period for request in ms since epoch
		Id string `json:"id"`
	}
)

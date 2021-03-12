package api

// swagger:model ErrorModel
type ErrorModel struct {
	// Service error code.
	// required: true
	Code ApiError `json:"code"`

	// Error hint message
	Msg string `json:"msg"`

	// Detailed error hint message
	DevMsg string `json:"devMsg"`

	// Request correlation id
	CorrId string `json:"corrId"`
}

// swagger:model Model
type Model struct {
	// The error response of the API
	//
	// required: true
	Error ErrorModel `json:"err"`

	// Business response of the service. Skipped if error occurs
	Data interface{} `json:"data,omitempty"`
}

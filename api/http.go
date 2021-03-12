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

type Http interface {
	AddGroup(name string, subPath string, config ConfigGroup) error
	AddRoute(method string, path string, config ConfigRoute, service Service) error
	AddGroupRoute(method string, path string, group string, config ConfigRoute, service Service)
	Listen(port int, address string)
}
package api

type ApiError int

const (
	ApiErrorNoError              ApiError = 0x0
	ApiErrorEntityDoesNotExists  ApiError = 0x1
	ApiErrorUnexpected           ApiError = 0x2
	ApiErrorEntityAlreadyExists  ApiError = 0x3
	ApiErrorUnknownItemRequested ApiError = 0x4
	ApiErrorMissingRequiredItem  ApiError = 0x5
	ApiErrorAuthFailed           ApiError = 0x6
)

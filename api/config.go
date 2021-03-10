package api

import "github.com/hellcats88/abstracte/tenant"

type HandlerReturn struct {
	HttpStatus    int
	Err           ApiError
	ErrNative     error
	ResponseModel interface{}
}

type RetriveTenant func(ctx interface{}) (tenant.Context, HandlerReturn)

type ApiConfig struct {
	TenantRetriever RetriveTenant
}

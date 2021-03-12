package api

import (
	"github.com/hellcats88/abstracte/runtime"
	"github.com/hellcats88/abstracte/tenant"
)

type ServiceInput interface {
	TenantCtx() tenant.Context
}

type ServiceOutput interface {
	Status() ApiError
	Err() error
	ErrMessage() string
	ResponseModel() interface{}
}

type Service func(runtime.Context, ServiceInput) (ServiceOutput, error)

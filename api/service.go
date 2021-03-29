package api

import (
	"github.com/hellcats88/abstracte/runtime"
)

type ServiceInput interface {
	Raw() interface{}
	RuntimeCtx() runtime.Context
	Model() interface{}
	InputParams() map[string]string
	QueryParams() interface{}
}

type ServiceOutput interface {
	Status() ApiError
	Err() error
	ErrMessage() string
	ResponseModel() interface{}
}

type Service func(runtime.Context, ServiceInput) (ServiceOutput, error)

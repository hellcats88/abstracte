package api

import (
	"github.com/hellcats88/abstracte/runtime"
)

type ServiceInput interface {
	RawCtx() interface{}
	Model() interface{}
	InputParams() map[string]string
	QueryParams() interface{}
	Headers() interface{}
}

type ServiceOutput interface {
	Status() ApiError
	Err() error
	ErrMessage() string
	ResponseModel() interface{}
}

type Service func(runtime.Context, ServiceInput) ServiceOutput

const ServiceResultKey string = "_abstracte_serviceresult_key"

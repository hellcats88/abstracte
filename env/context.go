package env

type Context interface {
	Name() string
}

type ContextFactory interface {
	Create() Context
}

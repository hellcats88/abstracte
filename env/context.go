package env

import "os"

type Context struct {
	name string
}

func New(name string) Context {
	return Context{name: name}
}

func NewFromEnvVar(name string) Context {
	return Context{name: os.Getenv(name)}
}

func (c Context) Name() string {
	return c.name
}

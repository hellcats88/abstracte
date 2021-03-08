package env

import "os"

type Context struct {
	Name string
}

func New(name string) Context {
	return Context{Name: name}
}

func NewFromEnvVar(name string) Context {
	return Context{Name: os.Getenv(name)}
}

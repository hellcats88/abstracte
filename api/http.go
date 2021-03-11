package api

type HttpGroup struct {
	Name    string
	SubPath string
}

type Http interface {
	AddGroup(item HttpGroup)
	AddRoute(path string)
	AddGroupRoute(path string, group string)
	Listen(port int, address string)
}

package api

type Http interface {
	AddGroup(name string)
	AddRoute(path string)
	AddGroupRoute(path string, group string)
	Listen(port int, address string)
}

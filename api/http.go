package api

type Http interface {
	AddGroup(name string, subPath string, config ApiConfigGroup) error
	AddRoute(path string)
	AddGroupRoute(path string, group string, config ApiConfigRoute)
	Listen(port int, address string)
}

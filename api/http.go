package api

type Http interface {
	AddGroup(name string, subPath string, config ApiConfigGroup) error
	AddRoute(path string)
	AddGroupRoute(path string, group string)
	Listen(port int, address string)
}

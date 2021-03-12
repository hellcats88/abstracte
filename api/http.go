package api

type Http interface {
	AddGroup(name string, subPath string, config ConfigGroup) error
	AddRoute(method string, path string, config ConfigRoute, service Service) error
	AddGroupRoute(method string, path string, group string, config ConfigRoute, service Service)
	Listen(port int, address string)
}

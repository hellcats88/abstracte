package api

type Http interface {
	AddGroup(name string, subPath string, config ConfigGroup) error
	AddRoute(path string, config ConfigRoute, service Service) error
	AddGroupRoute(path string, group string, config ConfigRoute, service Service)
	Listen(port int, address string)
}

package api

type Http interface {
	AddGroup(name string, subPath string, config Config) error
	AddRoute(method string, path string, config Config, service Service) error
	AddGroupRoute(method string, path string, group string, config Config, service Service) error
	Listen(port int, address string) error
}

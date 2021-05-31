package api

type Http interface {
	AddRoute(method string, path string, config Config, service Service) error
	Listen(port int, address string) error
}

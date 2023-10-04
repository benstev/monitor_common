package controller

// Controller is an interface for HTTP controllers
type Controller interface {
	DefineRoutes() error
	DefinePublicRoutes() error
}

package router

import (
	"github.com/labstack/echo/v4"
)

type Route struct {
	Method  string
	Path    string
	Handler echo.HandlerFunc
	// Roles []string
}

func PublicRoutes() []*Route {
	return []*Route{}
}

func privateRoutes() []*Route {
	return []*Route{}
}

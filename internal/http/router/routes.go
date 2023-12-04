package router

import (
	"github.com/fidya02/Capstone-Project/internal/http/handler"
	"github.com/labstack/echo/v4"
)

const (
	Administrator = "Administrator"
	Buyer         = "Buyer"
)

var (
	allRoles  = []string{Administrator, Buyer}
	onlyAdmin = []string{Administrator}
	onlyBuyer = []string{Buyer}
)

type Route struct {
	Method  string
	Path    string
	Handler echo.HandlerFunc
	Roles   []string
}

func PublicRoutes(authHandler *handler.AuthHandler) []*Route {
	return []*Route{
		{
			Method:  echo.POST,
			Path:    "/login",
			Handler: authHandler.Login,
		},
		{
			Method:  echo.POST,
			Path:    "/register",
			Handler: authHandler.Register,
		},
	}
}

func PrivateRoutes(userHandler *handler.UserHandler) []*Route {
	return []*Route{
		{
			Method:  echo.GET,
			Path:    "/users",
			Handler: userHandler.GetAllUsers,
			Roles:   allRoles,
		},
		{
			Method:  echo.GET,
			Path:    "/users/:id",
			Handler: userHandler.GetUserByID,
			Roles:   allRoles,
		},
		{
			Method:  echo.POST,
			Path:    "/users",
			Handler: userHandler.CreateUser,
			Roles:   onlyBuyer,
		},
		{
			Method:  echo.PUT,
			Path:    "/users/:id",
			Handler: userHandler.UpdateUser,
			Roles:   onlyBuyer,
		},
		{
			Method:  echo.DELETE,
			Path:    "/users/:id",
			Handler: userHandler.DeleteUser,
			Roles:   onlyAdmin,
		},
	}
}

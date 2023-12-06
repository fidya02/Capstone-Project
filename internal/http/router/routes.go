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

func PublicRoutes(
	authHandler *handler.AuthHandler,
	TicketHandler *handler.TicketHandler) []*Route {
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

		//Ticket
		{
			Method:  echo.GET,
			Path:    "/tickets",
			Handler: TicketHandler.FindAllTickets,
		},
		{
			Method:  echo.GET,
			Path:    "/tickets/:id",
			Handler: TicketHandler.GetTicketByID,
		},
		{
			Method:  echo.POST,
			Path:    "/tickets/category/:category",
			Handler: TicketHandler.FilterTicketByCategory,
		},
		{
			Method:  echo.POST,
			Path:    "/tickets/category",
			Handler: TicketHandler.FilterTicketByCategory,
		},
		{
			Method:  echo.GET,
			Path:    "/tickets/range/:min/:max",
			Handler: TicketHandler.FilterTicketByPriceRange,
		},
		{
			Method:  echo.GET,
			Path:    "/tickets/location/:location",
			Handler: TicketHandler.FilterTicketByLocation,
		},
		{
			Method:  echo.POST,
			Path:    "/tickets/available/:sort",
			Handler: TicketHandler.SortTicketByAvailable,
		},
		{
			Method:  echo.POST,
			Path:    "/tickets/available",
			Handler: TicketHandler.SortTicketByAvailable,
		},
		{
			Method:  echo.POST,
			Path:    "/tickets/sold",
			Handler: TicketHandler.SortTicketBySold,
		},
		{
			Method:  echo.POST,
			Path:    "/tickets/oldest",
			Handler: TicketHandler.SortTicketByOldest,
		},
		{

			Method:  echo.POST,
			Path:    "/tickets/newest",
			Handler: TicketHandler.SortTicketByNewest,
		},
		{
			Method:  echo.POST,
			Path:    "/tickets/cheap",
			Handler: TicketHandler.FilterTicketByCheap,
		},
		{
			Method:  echo.POST,
			Path:    "/tickets/expensive",
			Handler: TicketHandler.FilterTicketByExpensive,
		},
	}
}

func PrivateRoutes(
	userHandler *handler.UserHandler,
	TicketHandler *handler.TicketHandler,
	OrderHandler *handler.OrderHandler) []*Route {
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

		//Ticket
		{
			Method:  echo.GET,
			Path:    "/tickets",
			Handler: TicketHandler.FindAllTickets,
			Roles:   onlyAdmin,
		},
		{
			Method:  echo.POST,
			Path:    "/tickets",
			Handler: TicketHandler.CreateTicket,
			Roles:   onlyAdmin,
		},
		{
			Method:  echo.PUT,
			Path:    "/tickets/:id",
			Handler: TicketHandler.UpdateTicket,
			Roles:   onlyAdmin,
		},
		{
			Method:  echo.DELETE,
			Path:    "/tickets/:id",
			Handler: TicketHandler.DeleteTicket,
			Roles:   onlyAdmin,
		},
		{
			Method:  echo.POST,
			Path:    "/order",
			Handler: OrderHandler.CreateOrder,
			Roles:   allRoles,
		},

		{
			Method:  echo.GET,
			Path:    "/order",
			Handler: OrderHandler.GetAllOrders,
			Roles:   onlyAdmin,
		},

		{
			Method:  echo.GET,
			Path:    "/order/:id",
			Handler: OrderHandler.GetOrderByUserID,
			Roles:   allRoles,
		},
		//UserCreateOrder
		{
			Method:  echo.POST,
			Path:    "user/order",
			Handler: OrderHandler.UserCreateOrder,
			Roles:   onlyBuyer,
		},
		//GetOrderHistory
		{
			Method:  echo.GET,
			Path:    "user/order",
			Handler: OrderHandler.GetOrderHistory,
			Roles:   onlyBuyer,
		},
	}
}

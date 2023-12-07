package builder

import (
	"github.com/fidya02/Capstone-Project/internal/config"
	"github.com/fidya02/Capstone-Project/internal/http/handler"
	"github.com/fidya02/Capstone-Project/internal/http/router"
	"github.com/fidya02/Capstone-Project/internal/repository"
	"github.com/fidya02/Capstone-Project/internal/service"

	"gorm.io/gorm"
)

func BuildPublicRoutes(cfg *config.Config, db *gorm.DB) []*router.Route {
	registerRepository := repository.NewRegisterRepository(db)
	registerService := service.NewRegisterService(registerRepository)
	userRepository := repository.NewUserRepository(db)
	loginService := service.NewLoginService(userRepository)
	tokenService := service.NewTokenService(cfg)
	BlogRepository := repository.NewBlogRepository(db)
	BlogService := service.NewBlogService(BlogRepository)
	BlogHandler := handler.NewBlogHandler(BlogService)
	ticketRepository := repository.NewTicketRepository(db)
	ticketService := service.NewTicketRepository(ticketRepository)
	tickethandler := handler.NewTicketHandler(ticketService)
	authHandler := handler.NewAuthHandler(registerService, loginService, tokenService)
	return router.PublicRoutes(authHandler, tickethandler, BlogHandler)
}

func BuildPrivateRoutes(cfg *config.Config, db *gorm.DB) []*router.Route {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	ticketRepository := repository.NewTicketRepository(db)
	ticketService := service.NewTicketRepository(ticketRepository)
	tickethandler := handler.NewTicketHandler(ticketService)
	userHandler := handler.NewUserHandler(cfg, userService)

	BlogRepository := repository.NewBlogRepository(db)
	BlogService := service.NewBlogService(BlogRepository)
	BlogHandler := handler.NewBlogHandler(BlogService)

	NotificationRepository := repository.NewNotificationRepository(db)
	NotificationService := service.NewNotificationService(NotificationRepository)
	NotificationHandler := handler.NewNotificationHandler(NotificationService)

	OrderRepository := repository.NewOrderRepository(db)
	OrderService := service.NewOrderService(OrderRepository)
	OrderHandler := handler.NewOrderHandler(OrderService)

	return router.PrivateRoutes(userHandler, OrderHandler, BlogHandler, tickethandler, OrderHandler, NotificationHandler)

}

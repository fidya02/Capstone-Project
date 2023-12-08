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

	transactionRepository := repository.NewTransactionRepository(db)
	paymentService := service.NewPaymentService(midtransClient)
	transactionService := service.NewTransactionService(transactionRepository)
	transactionHandler := handler.NewTransactionHandler(transactionService, paymentService)

	loginService := service.NewLoginService(userRepository)
	tokenService := service.NewTokenService(cfg)
	authHandler := handler.NewAuthHandler(registerService, loginService, tokenService)
	
	BlogRepository := repository.NewBlogRepository(db)
	BlogService := service.NewBlogService(BlogRepository)
	BlogHandler := handler.NewBlogHandler(BlogService)
	
	ticketRepository := repository.NewTicketRepository(db)
	ticketService := service.NewTicketRepository(ticketRepository)
	tickethandler := handler.NewTicketHandler(ticketService)

	return router.PublicRoutes(authHandler, tickethandler, BlogHandler, transactionHandler)
}

func BuildPrivateRoutes(cfg *config.Config, db *gorm.DB) []*router.Route {
	userRepository := repository.NewUserRepository(db)
	ticketRepository := repository.NewTicketRepository(db)
	ticketService := service.NewTicketRepository(ticketRepository)
	tickethandler := handler.NewTicketHandler(ticketService)
	userService := service.NewUserService(userRepository)
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

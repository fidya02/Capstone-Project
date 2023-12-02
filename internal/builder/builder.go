package builder

import (
	"github.com/fidya02/Capstone-Project/internal/config"
	"github.com/fidya02/Capstone-Project/internal/http/router"

	"gorm.io/gorm"
)

func BuildPublicRoutes(cfg *config.Config, db *gorm.DB) []*router.Route {
	// userRepository := repository.NewUserRepository(db)
	// loginService := service.NewLoginService(userRepository)
	// tokenService := service.NewTokenService(cfg)
	// authHandler := handler.NewAuthHandler(loginService, tokenService)
	return router.PublicRoutes()
	// return router.PublicRoutes(authHandler)
}

func BuildPrivateRoutes(cfg *config.Config, db *gorm.DB) []*router.Route {
	// userRepository := repository.NewUserRepository(db)
	// userService := service.NewUserService(userRepository)
	// userHandler := handler.NewUserHandler(cfg, userService)
	return router.PrivateRoutes()
	// return router.PrivateRoutes(userHandler)
}

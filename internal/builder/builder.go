package builder

import (
	"GOLANG/internal/config"
	"GOLANG/internal/http/handler"
	"GOLANG/internal/http/router"
	"GOLANG/internal/repository"
	"GOLANG/internal/service"

	"gorm.io/gorm"
)

func BuildPublicRoutes(cfg *config.Config, db *gorm.DB) []*router.Route {
	userRepository := repository.NewUserRepository(db) // kenapa make ini? karena nge find email nya dari user_repository
	loginService := service.NewLoginService(userRepository)
	tokenService := service.NewTokenService(cfg)
	authHandler := handler.NewAuthHandler(loginService, tokenService)
	return router.PublicRoutes(authHandler)
}

func BuildPrivateRoutes(cfg *config.Config, db *gorm.DB) []*router.Route {
	//memanggil fungsi PublicRoutes() dari router
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	return router.PrivateRoutes(userHandler)
}

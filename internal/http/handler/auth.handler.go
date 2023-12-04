package handler

import (
	"net/http"

	"github.com/fidya02/Capstone-Project/entity"
	"github.com/fidya02/Capstone-Project/internal/http/validator"
	"github.com/fidya02/Capstone-Project/internal/service"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	loginService    service.LoginUseCase
	registerService service.RegisterUseCase
	tokenService    service.TokenUseCase
}

func NewAuthHandler(
	loginService service.LoginUseCase,
	registerService service.RegisterUseCase,
	tokenService service.TokenUseCase,
) *AuthHandler {
	return &AuthHandler{
		loginService:    loginService,
		registerService: registerService,
		tokenService:    tokenService,
	}
}

func (h *AuthHandler) Login(ctx echo.Context) error {
	var input struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	user, err := h.loginService.Login(ctx.Request().Context(), input.Email, input.Password)

	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	accessToken, err := h.tokenService.GenerateAccessToken(ctx.Request().Context(), user)

	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	data := map[string]string{
		"access_token": accessToken,
	}

	return ctx.JSON(http.StatusOK, data)
}

func (h *AuthHandler) Regist(ctx echo.Context) error {
	var input struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
		Roles    string `json:"roles" default:"Buyer"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	user := entity.Regist(input.Email, input.Password, input.Roles)
	err := h.registerService.Registration(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	accessToken, err := h.tokenService.GenerateAccessToken(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message":      "User register successfully",
		"access_token": accessToken,
	})
}

package handler

import (
	"net/http"

	"github.com/fidya02/Capstone-Project/entity"
	"github.com/fidya02/Capstone-Project/internal/config"
	"github.com/fidya02/Capstone-Project/internal/http/validator"
	"github.com/fidya02/Capstone-Project/internal/service"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	cfg         *config.Config
	userService service.UserUseCase
}

func NewUserHandler(cfg *config.Config, userService service.UserUseCase) *UserHandler {
	return &UserHandler{cfg, userService}
}

func (h *UserHandler) GetAllUsers(ctx echo.Context) error {
	users, err := h.userService.GetAll(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	return ctx.JSON(http.StatusOK, users)
}

func (h *UserHandler) CreateUser(ctx echo.Context) error {
	var input struct {
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"required"`
		Number   string `json:"number" validate:"required"`
		Wallet   int64  `json:"wallet"`
		Password string `json:"password" validate:"required"`
		Role     string `json:"role" validate:"required,oneof=Administrator Buyer"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	user := entity.NewUser(input.Name, input.Number, input.Email, input.Password, input.Role, input.Wallet)
	err := h.userService.CreateUser(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusCreated, user)
}

func (h *UserHandler) UpdateUser(ctx echo.Context) error {
	var input struct {
		ID       int64  `param:"id"`
		Name     string `json:"name"`
		Wallet   int64  `json:"wallet"`
		Number   string `json:"number"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	user := entity.UpdateUser(input.ID, input.Wallet, input.Number, input.Name, input.Email, input.Password, input.Role)

	err := h.userService.UpdateUser(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]string{"success": "succesfully update user"})
}

func (h *UserHandler) DeleteUser(ctx echo.Context) error {
	var input struct {
		ID int64 `param:"id" validate:"required"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	err := h.userService.Delete(ctx.Request().Context(), input.ID)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (h *UserHandler) GetUserByID(ctx echo.Context) error {
	var input struct {
		ID int64 `param:"id" validate:"required"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	user, err := h.userService.FindByID(ctx.Request().Context(), input.ID)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, user)
}

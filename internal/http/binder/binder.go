package binder

import (
	"github.com/creasty/defaults"
	internalValidator "github.com/fidya02/Capstone-Project/internal/http/validator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Binder struct {
	defaultBinder *echo.DefaultBinder
	*internalValidator.FormValidator
}

func NewBinder(
	dbr *echo.DefaultBinder,
	vdr *internalValidator.FormValidator) *Binder {
	return &Binder{dbr, vdr}
}

func (b *Binder) Bind(i interface{}, c echo.Context) error {
	if err := b.defaultBinder.Bind(i, c); err != nil {
		return err
	}

	if err := defaults.Set(i); err != nil {
		return err
	}

	if err := b.Validate(i); err != nil {
		errs := err.(validator.ValidationErrors)
		return errs
	}

	return nil
}

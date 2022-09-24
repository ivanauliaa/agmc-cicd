package middleware

import (
	"day9-cicd/pkg/util/validator"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	LogMiddleware(e)

	e.Validator = &validator.CustomValidator{
		Validator: validator.NewValidator(),
	}
}

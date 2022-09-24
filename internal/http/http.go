package http

import (
	"day9-cicd/internal/app/auth"
	"day9-cicd/internal/app/book"
	"day9-cicd/internal/app/user"
	"day9-cicd/internal/factory"

	"github.com/labstack/echo/v4"
)

func NewHttp(e *echo.Echo, f *factory.Factory) {
	v1G := e.Group("/v1")

	user.NewHandler(f).Route(v1G.Group("/users"))
	auth.NewHandler(f).Route(v1G.Group("/auth"))
	book.NewHandler(f).Route(v1G.Group("/books"))
}

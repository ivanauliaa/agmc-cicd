package http_test

import (
	"day9-cicd/database/migration"
	"day9-cicd/internal/factory"
	server "day9-cicd/internal/http"
	"day9-cicd/internal/middleware"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	os.Setenv("APP_ENV", "test")
	godotenv.Load("../../.env")
	migration.Migrate()
}

func TestHello_Hello(t *testing.T) {
	f := factory.NewFactory()
	e := echo.New()
	middleware.Init(e)
	server.NewHttp(e, f)

	server := httptest.NewServer(e)
	defer server.Close()

	ex := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  server.URL,
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	t.Run("should return 200", func(t *testing.T) {
		resBody := ex.GET("/v1/hello").
			Expect().
			Status(http.StatusOK).
			JSON().Object()

		resBody.Value("message").Equal("hello world")
	})
}

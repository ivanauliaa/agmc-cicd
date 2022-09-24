package user_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"day9-cicd/database/migration"
	"day9-cicd/internal/factory"
	server "day9-cicd/internal/http"
	"day9-cicd/internal/middleware"
	"day9-cicd/pkg/helper"

	"github.com/gavv/httpexpect/v2"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	os.Setenv("APP_ENV", "test")
	godotenv.Load("../../../.env")
	migration.Migrate()
}

func TestHandler_Create(t *testing.T) {
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

	t.Run("should return 201 if meet payload requirements", func(t *testing.T) {
		t.Cleanup(func() {
			helper.CleanUsersTable()
		})

		resBody := ex.POST("/v1/users").
			WithJSON(map[string]interface{}{
				"name":     "agmc",
				"email":    "agmc@gmail.com",
				"password": "agmc",
			}).
			Expect().
			Status(http.StatusCreated).
			JSON().Object()

		resBody.Value("status").Equal("success")
	})
}

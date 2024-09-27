package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/rajdeep-ray/glofox/handlers"
	"github.com/stretchr/testify/assert"
)

func TestGetClasses(t *testing.T) {
    app := fiber.New()
    app.Get("/classes", handlers.ListClasses)

    req := httptest.NewRequest("GET", "/classes", nil)
    resp, _ := app.Test(req)

    assert.Equal(t, 200, resp.StatusCode)
}

func TestCreateClass(t *testing.T) {
    app := fiber.New()
    app.Post("/classes", handlers.AddClasses)

    req := httptest.NewRequest("POST", "/classes", nil)
    req.Header.Set("Content-Type", "application/json")
    resp, _ := app.Test(req)

    assert.Equal(t, 400, resp.StatusCode)
}

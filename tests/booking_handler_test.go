package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/rajdeep-ray/glofox/handlers"
	"github.com/stretchr/testify/assert"
)

func TestGetBookings(t *testing.T) {
    app := fiber.New()
    app.Get("/booking", handlers.ListBookings)

    req := httptest.NewRequest("GET", "/booking", nil)
    resp, _ := app.Test(req)

    assert.Equal(t, 200, resp.StatusCode)
}

func TestCreateBooking(t *testing.T) {
    app := fiber.New()
    app.Post("/booking", handlers.AddBooking)

    req := httptest.NewRequest("POST", "/booking", nil)
    req.Header.Set("Content-Type", "application/json")
    resp, _ := app.Test(req)

    assert.Equal(t, 400, resp.StatusCode)
}

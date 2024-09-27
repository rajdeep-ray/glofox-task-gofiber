package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rajdeep-ray/glofox/handlers"
)

func SetupRoutes(app *fiber.App) {
    app.Get("/classes", handlers.ListClasses)
    app.Post("/classes", handlers.AddClasses)
    
    app.Get("/booking", handlers.ListBookings)
    app.Post("/booking", handlers.AddBooking)
}

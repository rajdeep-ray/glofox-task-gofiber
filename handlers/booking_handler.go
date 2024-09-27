package handlers

import (
	"fmt"

	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rajdeep-ray/glofox/models"
	"github.com/rajdeep-ray/glofox/requests"
)

var bookings = []models.Booking{}

func matchingClass(className string, date time.Time) *models.Class {
	for _, class := range classes {
		if class.Name == className && class.Date.Equal(date) {
			return &class
		}
	}
	return nil
}

func AddBooking(c *fiber.Ctx) error {
	var reqBooking requests.BookingRequest

	if err := c.BodyParser(&reqBooking); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if reqBooking.Name == "" || reqBooking.Date == "" || reqBooking.ClassName == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid booking details"})
	}

	const timeLayout = "2006-01-02"
	date, err := time.Parse(timeLayout, reqBooking.Date)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid date input",
		})
	}

	class := matchingClass(reqBooking.ClassName, date)
	if class == nil {
		return c.Status(404).JSON(fiber.Map{"error": "No Class Found!"})
	}

	bookings = append(bookings, models.Booking{
		Id:    len(bookings) + 1,
		Name:  reqBooking.Name,
		Date:  date,
		Class: class,
	})

	return c.Status(201).JSON(fiber.Map{
		"message": "Booking Done!",
	})
}

func ListBookings(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message":  fmt.Sprintf("%v Bookings found!", len(bookings)),
		"bookings": bookings,
	})
}

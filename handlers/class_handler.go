package handlers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rajdeep-ray/glofox/models"
	"github.com/rajdeep-ray/glofox/requests"
)

var classes = []models.Class{}

func ListClasses(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": fmt.Sprintf("%v classes found!", len(classes)),
		"classes": classes,
	})
}

func AddClasses(c *fiber.Ctx) error {
	var reqClass requests.ClassRequest

	if err := c.BodyParser(&reqClass); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if reqClass.Name == "" || reqClass.StartDate == "" || reqClass.EndDate == "" || reqClass.Capacity <= 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid class details"})
	}

	const timeLayout = "2006-01-02"
	startDate, err := time.Parse(timeLayout, reqClass.StartDate)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid startDate input",
		})
	}

	endDate, err := time.Parse(timeLayout, reqClass.EndDate)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid endDate input",
		})
	}

	for d, i := startDate, 1; !d.After(endDate); d, i = d.AddDate(0, 0, 1), i+1 {
		classes = append(classes, models.Class{
			Id:       i,
			Name:     reqClass.Name,
			Capacity: reqClass.Capacity,
			Date:     d,
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Classes added",
	})
}

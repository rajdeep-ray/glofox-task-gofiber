package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rajdeep-ray/glofox/routes"
)

func main() {
	app := fiber.New()

	routes.SetupRoutes(app)

	app.Listen(":3000")
}

// type Class struct {
// 	Id       int       `json:"id"`
// 	Name     string    `json:"name"`
// 	Date     time.Time `json:"startDate"`
// 	Capacity int       `json:"capacity"`
// }

// type Booking struct {
// 	Id    int       `json:"id"`
// 	Name  string    `json:"name"`
// 	Date  time.Time `json:"date"`
// 	Class *Class    `json:"class"`
// }

// type ClassRequest struct {
// 	Name      string `json:"name"`
// 	StartDate string `json:"startDate"`
// 	EndDate   string `json:"endDate"`
// 	Capacity  int    `json:"capacity"`
// }

// type BookingRequest struct {
// 	Name      string `json:"name"`
// 	Date      string `json:"date"`
// 	ClassName string `json:"class"`
// }

// var classes = []Class{}

// var bookings = []Booking{}

// func MatchingClass(className string, date time.Time) *Class {
// 	for _, class := range classes {
// 		if class.Name == className && class.Date.Equal(date) {
// 			return &class
// 		}
// 	}
// 	return nil
// }

// func main() {

// 	app := fiber.New()

// 	app.Get("/", func(c *fiber.Ctx) error {
// 		return c.Status(200).JSON(fiber.Map{"message": "Hello, World!"})
// 	})

// 	// Lits all Classes
// 	app.Get("/classes", func(c *fiber.Ctx) error {
// 		return c.Status(200).JSON(fiber.Map{
// 			"message": fmt.Sprintf("%v classes found!", len(classes)),
// 			"classes": classes,
// 		})
// 	})

// 	app.Post("/classes", func(c *fiber.Ctx) error {
// 		var reqClass ClassRequest

// 		if err := c.BodyParser(&reqClass); err != nil {
// 			return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
// 		}

// 		// checks for valids req fields
// 		if reqClass.Name == "" || reqClass.StartDate == "" || reqClass.EndDate == "" || reqClass.Capacity <= 0 {
// 			return c.Status(400).JSON(fiber.Map{"error": "Invalid class details"})
// 		}

// 		const timeLayout = "2006-01-02"
// 		startDate, err := time.Parse(timeLayout, reqClass.StartDate)

// 		if err != nil {
// 			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 				"error": "Invalid startDate input",
// 			})
// 		}

// 		endDate, err := time.Parse(timeLayout, reqClass.EndDate)
// 		if err != nil {
// 			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 				"error": "Invalid endDate input",
// 			})
// 		}

// 		for d, i := startDate, 1; !d.After(endDate); d, i = d.AddDate(0, 0, 1), i+1 {
// 			classes = append(classes, Class{
// 				Id:       i,
// 				Name:     reqClass.Name,
// 				Capacity: reqClass.Capacity,
// 				Date:     d,
// 			})
// 		}

// 		return c.Status(201).JSON(fiber.Map{
// 			"message": "Classes added",
// 		})
// 	})

// 	app.Get("/booking", func(c *fiber.Ctx) error {

// 		return c.Status(200).JSON(fiber.Map{
// 			"message":  fmt.Sprintf("%v Bookings found!", len(bookings)),
// 			"bookings": bookings,
// 		})
// 	})

// 	app.Post("/booking", func(c *fiber.Ctx) error {
// 		var reqBooking BookingRequest

// 		if err := c.BodyParser(&reqBooking); err != nil {
// 			return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
// 		}

// 		// checks for valids req fields
// 		if reqBooking.Name == "" || reqBooking.Date == "" || reqBooking.ClassName == "" {
// 			return c.Status(400).JSON(fiber.Map{"error": "Invalid class details"})
// 		}

// 		const timeLayout = "2006-01-02"
// 		date, err := time.Parse(timeLayout, reqBooking.Date)

// 		if err != nil {
// 			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 				"error": "Invalid startDate input",
// 			})
// 		}

// 		class := MatchingClass(reqBooking.ClassName, date)

// 		if class == nil {
// 			return c.Status(404).JSON(fiber.Map{
// 				"error": "No Class Found!",
// 			})
// 		}

// 		bookings = append(bookings, Booking{
// 			Id:    len(bookings) + 1,
// 			Name:  reqBooking.Name,
// 			Date:  date,
// 			Class: class,
// 		})

// 		return c.Status(201).JSON(fiber.Map{
// 			"message": "Booking Done!",
// 		})
// 	})

// 	app.Listen(":3000")
// }

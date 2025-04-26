package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"smart-grocery-agent/internal/services" 
)

// GenerateGroceryList is a Fiber-compatible handler
func GenerateGroceryList(c *fiber.Ctx) error {
	// Parse the request body
	var request struct {
		Meals []string `json:"meals"`
	}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Call the GenerateGroceryList function from ai_agent.go
	groceryList, err := services.GenerateGroceryList(request.Meals)
	if err != nil {
		log.Printf("Error generating grocery list: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Return the grocery list as a JSON response
	return c.JSON(fiber.Map{
		"grocery_list": groceryList,
	})
}

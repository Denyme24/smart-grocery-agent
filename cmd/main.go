package main

import (
    "log"
    "os"

    "github.com/gofiber/fiber/v2"
    "github.com/joho/godotenv"
    "smart-grocery-agent/internal/handlers"
)

func main() {
    // Print the current working directory
    wd, _ := os.Getwd()
    log.Printf("Current working directory: %s", wd)

    // Load environment variables from the .env file
    if err := godotenv.Load(".env"); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // Print the GEMINI_API_KEY to confirm it is loaded
    apiKey := os.Getenv("GEMINI_API_KEY")
    if apiKey == "" {
        log.Fatalf("GEMINI_API_KEY environment variable is empty or not set")
    }
    log.Printf("GEMINI_API_KEY: %s", apiKey)

    app := fiber.New()

    // Define the route and handler
    app.Post("/grocery-list", handlers.GenerateGroceryList)

    log.Printf("Server starting on http://localhost:3000")
    // Start the server
    log.Fatal(app.Listen(":3000"))
}
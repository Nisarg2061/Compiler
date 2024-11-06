package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// Function to run the lexer and return the output as a string
func runLexer() (string, error) {
	cmd := exec.Command("go", "run", "./src/go_lexer/main.go")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error executing lexical analysis: %s", err)
	}
	return string(out), nil
}

// Function to run semantic analysis and return the output as a string
func runSemantic() (string, error) {
	cmd := exec.Command("python3", "./src/python_parser/semantic_analyzer.py")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error executing semantic analysis: %s", err)
	}
	return string(out), nil
}

// Function to run the 3AC generation and return the output as a string
func run3ac() (string, error) {
	cmd := exec.Command("python3", "./src/3ac/3ac.py")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error executing 3AC generation: %s", err)
	}
	return string(out), nil
}

// Route to handle file uploads
func uploadFile(c *fiber.Ctx) error {
	// Get the file from the form data
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "File is required",
		})
	}

	// Create or open the "sample.txt" file
	dst := "./sample.txt"
	if err := c.SaveFile(file, dst); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to save file: %s", err.Error()),
		})
	}

	// Return a success message
	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("File uploaded successfully as %s", dst),
	})
}

// Main function to set up the Fiber app and routes
func main() {
	// Initialize a new Fiber app
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Or specify "http://frontend:5099"
		AllowMethods: "POST, GET, OPTIONS",
	}))

	// Route for lexer
	app.Get("/lexer", func(c *fiber.Ctx) error {
		output, err := runLexer()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"output": output,
		})
	})

	// Route for semantic analysis
	app.Get("/semantic", func(c *fiber.Ctx) error {
		output, err := runSemantic()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"output": output,
		})
	})

	// Route for 3AC generation
	app.Get("/3ac", func(c *fiber.Ctx) error {
		output, err := run3ac()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"output": output,
		})
	})

	// Route for running all tasks (lexer, semantic, 3ac)
	app.Get("/all", func(c *fiber.Ctx) error {
		lexerOutput, err := runLexer()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		semanticOutput, err := runSemantic()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		threeACOutput, err := run3ac()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"lexer":    lexerOutput,
			"semantic": semanticOutput,
			"3ac":      threeACOutput,
		})
	})

	// Route for file upload
	app.Post("/upload", uploadFile)

	// Start the server
	log.Fatal(app.Listen(":3000"))
}

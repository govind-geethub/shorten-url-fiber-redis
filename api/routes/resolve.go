package routes

import "github.com/gofiber/fiber/v2"

// ResolveURL handles redirecting short URLs back to the original long URLs
func ResolveURL(c *fiber.Ctx) error {
	return c.SendString("Hello from the resolve endpoint!")
}

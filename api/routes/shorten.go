package routes

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	helpers "github.com/govind-geethub/shorten-url-fiber-redis/helper"
)

type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

type response struct {
	URL            string        `json:"url"`
	CustomShort    string        `json:"short"`
	Expiry         time.Duration `json:"expiry"`
	XRateRemaining int           `json:"rate_limit"`
	XRateLimitRest time.Duration `json:"rate_limit_reset"`
}

// ShortenURL handles creating a shortened string for a given long URL
func ShortenURL(c *fiber.Ctx) error {

	body := new(request)

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// implement rate limiting
	// check the IP if exists then decrease its limit by 1

	// check if the input is the acutal url or not
	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid URL"})
	}

	// check for domain error
	// may abuse shortener which may lead to infinite loop

	if !helpers.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error": "You can't hack it😁"})
	}

	// enforce https, SSL
	body.URL = helpers.EnforceHTTP(body.URL)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "URL successfully processed",
		"url":     body.URL,
	})
}

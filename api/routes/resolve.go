package routes // named on the folder makes it easy to import and accessible

import (
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/govind-geethub/shorten-url-fiber-redis/database"
)

// ResolveURL handles redirecting short URLs back to the original long URLs
func ResolveURL(c *fiber.Ctx) error {

	url := c.Params("url")

	r := database.CreateClient(0)
	defer r.Close()

	value, err := r.Get(database.Ctx, url).Result()
	if err == redis.Nil { // couldn't find
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "short not found int he database",
		})
	} else if err != nil { // cannot connect
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot connect to DB",
		})
	}

	rInr := database.CreateClient(1)
	defer rInr.Close()

	_ = rInr.Incr(database.Ctx, "counter")

	return c.Redirect(value, 301) // with status 301 redirecting to the actual link
}

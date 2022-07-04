package posts_router

import (
	"github.com/gofiber/fiber/v2"
	postctrl "github.com/isaacismaelx14/blog-api/controller/posts"
)

func get(_route fiber.Router) {
	_route.Get("/", _get)
}

func _get(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(postctrl.Get())
}

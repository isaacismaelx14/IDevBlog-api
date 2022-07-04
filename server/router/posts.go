package router

import (
	"github.com/gofiber/fiber/v2"
	postctrl "github.com/isaacismaelx14/blog-api/controller/posts"
)

func posts() {
	group := app.Group("/posts")
	group.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(postctrl.Get())
	})
}

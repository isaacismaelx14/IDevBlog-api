package posts_router

import (
	"github.com/gofiber/fiber/v2"
	message_ctrl "github.com/isaacismaelx14/blog-api/controller/message"
	post_ctrl "github.com/isaacismaelx14/blog-api/controller/posts"
	"github.com/isaacismaelx14/blog-api/schemas"
)

func post(_route fiber.Router) {
	_route.Post("/", _post)
}

func _post(c *fiber.Ctx) error {
	// create post
	post := schemas.Post{}
	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(message_ctrl.Error(fiber.StatusBadRequest, err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(message_ctrl.Message(post_ctrl.Create(post)))
}

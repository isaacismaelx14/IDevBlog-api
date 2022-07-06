package users_router

import (
	"github.com/gofiber/fiber/v2"
	message_ctrl "github.com/isaacismaelx14/blog-api/controller/message"
	users_ctrl "github.com/isaacismaelx14/blog-api/controller/users"
	"github.com/isaacismaelx14/blog-api/schemas"
)

func post(_route fiber.Router) {
	_route.Post("/", _post)
}

func _post(c *fiber.Ctx) error {
	// create user
	user := schemas.User{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(message_ctrl.Error(fiber.StatusBadRequest, err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(message_ctrl.Message(users_ctrl.Create(user)))
}

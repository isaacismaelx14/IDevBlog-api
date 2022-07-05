package posts_router

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	message_ctrl "github.com/isaacismaelx14/blog-api/controller/message"
	post_ctrl "github.com/isaacismaelx14/blog-api/controller/posts"
)

func get(_route fiber.Router) {
	_route.Get("/", _get)
}

func _get(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(message_ctrl.Message(post_ctrl.Get()))
}

func getOne(_route fiber.Router) {
	_route.Get("/:id", _getOne)
}

func _getId(_id string) (uint, error) {
	id, err := strconv.ParseInt(_id, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

func _getOne(c *fiber.Ctx) error {
	id, err := _getId(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(message_ctrl.Error(fiber.StatusBadRequest, "Invalid id"))
	}

	post := post_ctrl.GetOne(uint(id))

	if post.Id == 0 {
		code := fiber.StatusNotFound
		return c.Status(code).JSON(message_ctrl.Error(code, "Post not found"))
	}

	return c.Status(fiber.StatusOK).JSON(message_ctrl.Message(post))
}

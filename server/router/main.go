package router

import (
	"github.com/gofiber/fiber/v2"
	posts_router "github.com/isaacismaelx14/blog-api/server/router/posts"
	users_router "github.com/isaacismaelx14/blog-api/server/router/users"
)

var app *fiber.App

func Init(_app *fiber.App) {
	app = _app
	home()
	posts_router.Init(app)
	users_router.Init(app)
}

package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

var app *fiber.App

func Init(_app *fiber.App) {
	app = _app
	_cors()
}

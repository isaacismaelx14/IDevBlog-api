package middlewares

import "github.com/gofiber/fiber/v2/middleware/cors"

func _cors() {
	app.Use(cors.New())
}

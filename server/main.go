package server

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/isaacismaelx14/blog-api/server/middlewares"
	"github.com/isaacismaelx14/blog-api/server/router"
)

var app = fiber.New()

func getPort(_port int) string {
	port := os.Getenv("PORT")
	if port == "" {
		port = strconv.Itoa(_port)
	}

	return fmt.Sprintf(":%v", port)
}

func Init(_port int) {
	middlewares.Init(app)
	router.Init(app)

	app.Listen(getPort(_port))
}

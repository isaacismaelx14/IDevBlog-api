package main

import (
	"github.com/isaacismaelx14/blog-api/config"
	"github.com/isaacismaelx14/blog-api/server"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	config.Init()
	server.Init(4000)
}

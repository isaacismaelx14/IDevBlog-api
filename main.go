package main

import (
	database_ctrl "github.com/isaacismaelx14/blog-api/controller/database"
	"github.com/isaacismaelx14/blog-api/server"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	database_ctrl.Init()
	server.Init(4000)
}

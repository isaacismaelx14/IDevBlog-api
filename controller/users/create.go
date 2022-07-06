package users_ctrl

import (
	"github.com/isaacismaelx14/blog-api/config"
	"github.com/isaacismaelx14/blog-api/schemas"
)

func Create(User schemas.User) string {
	User.Password = string(config.GeneratePass(User.Password))
	config.DataBase.Create(&User)
	return "User created"
}

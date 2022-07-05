package post_ctrl

import (
	database_ctrl "github.com/isaacismaelx14/blog-api/controller/database"
	"github.com/isaacismaelx14/blog-api/schemas"
)

func Get() []schemas.Post {
	var posts = []schemas.Post{}
	database_ctrl.DataBase.Find(&posts)
	return posts
}

func GetOne(id uint) schemas.Post {
	var post = schemas.Post{}
	request := database_ctrl.DataBase.Where("id = ?", id).Find(&post)
	if request.RowsAffected == 0 {
		return schemas.Post{}
	}
	return post
}

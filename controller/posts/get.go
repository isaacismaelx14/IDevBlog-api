package post_ctrl

import (
	"github.com/isaacismaelx14/blog-api/config"
	"github.com/isaacismaelx14/blog-api/schemas"
)

func Get() []schemas.Post {
	var posts = []schemas.Post{}
	config.DataBase.Find(&posts)
	return posts
}

func GetOne(id uint) schemas.Post {
	var post = schemas.Post{}
	request := config.DataBase.Where("id = ?", id).Find(&post)
	if request.RowsAffected == 0 {
		return schemas.Post{}
	}
	return post
}

package post_ctrl

import (
	database_ctrl "github.com/isaacismaelx14/blog-api/controller/database"
	message_ctrl "github.com/isaacismaelx14/blog-api/controller/message"
	"github.com/isaacismaelx14/blog-api/schemas"
)

func Create(post schemas.Post) message_ctrl.IMessage {
	database_ctrl.DataBase.Create(&post)
	return message_ctrl.Message(post)
}

package post_ctrl

import (
	"github.com/isaacismaelx14/blog-api/config"
	message_ctrl "github.com/isaacismaelx14/blog-api/controller/message"
	"github.com/isaacismaelx14/blog-api/schemas"
)

func Create(post schemas.Post) message_ctrl.IMessage {
	config.DataBase.Create(&post)
	return message_ctrl.Message(post)
}

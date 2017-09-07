package views

import "webrtc-china.org/models"
import "time"

type ReplyView struct {
	TopicId   int64     `json:"topic_id"`
	UserView  *UserView `json:"user"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func BuildReplyView(topicId int64, reply *models.Reply, user *models.User) ReplyView {
	userView := BuildUserView(user)
	replyView := ReplyView{
		TopicId:   topicId,
		Content:   reply.Content,
		UserView:  &userView,
		CreatedAt: reply.CreatedAt,
	}
	return replyView
}

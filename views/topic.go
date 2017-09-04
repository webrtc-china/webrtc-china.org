package views

import "webrtc-china.org/models"

type TopicView struct {
	TopicId  int64     `json:"topic_id"`
	Title    string    `json:"title"`
	UserView *UserView `json:"user"`
	Content  string    `json:"content"`
	Node     string    `json:"node"`
}

func BuildTopicView(topic *models.Topic, user *models.User) TopicView {
	userView := BuildUserView(user)
	topicView := TopicView{
		TopicId:  topic.Id,
		Title:    topic.Title,
		Content:  topic.Content,
		Node:     topic.Node,
		UserView: &userView,
	}
	return topicView
}

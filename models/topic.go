package models

import (
	"context"
	"log"
	"time"

	"webrtc-china.org/session"
)

type Topic struct {
	Id        int64
	Title     string
	Content   string
	UserId    string
	Node      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CreateTopic(ctx context.Context, user *User, title string, content string, node string) (*Topic, error) {
	topic := &Topic{
		Title:     title,
		Content:   content,
		UserId:    user.Id,
		Node:      node,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	err := session.Database(ctx).Insert(topic)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return topic, nil
}

func GetAllTopics() {

}

func FindTopicById(id int64) {

}

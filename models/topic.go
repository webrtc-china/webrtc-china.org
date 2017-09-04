package models

import (
	"time"

	"github.com/go-pg/pg"
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

func CreateTopic(db *pg.DB, userId string, title string, content string, node string) (*Topic, error) {
	topic := &Topic{
		Title:   title,
		Content: content,
		UserId:  userId,
		Node:    node,
	}
	err := db.Insert(topic)
	if err != nil {
		return nil, err
	}
	return topic, nil
}

func GetAllTopics() {

}

func FindTopicById(id int64) {

}

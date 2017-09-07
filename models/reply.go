package models

import (
	"context"
	"time"

	"webrtc-china.org/session"
)

type Reply struct {
	Id        int64
	Content   string
	UserId    string
	TopicId   int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CreateReply(ctx context.Context, content string, userId string, topicId int64) (*Reply, error) {
	reply := &Reply{
		Content:   content,
		UserId:    userId,
		TopicId:   topicId,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	err := session.Database(ctx).Insert(reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

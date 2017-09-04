package models

import (
	"time"
)

type Reply struct {
	Id        int64
	Content   string
	UserId    string
	TopicId   string
	Node      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

package model

import "time"

type MessageDetail struct {
	TimeStamp time.Time
	User      string
	Message   string
}

func NewMessageDetail(timestamp time.Time, user string, message string) *MessageDetail {
	return &MessageDetail{TimeStamp: timestamp, User: user, Message: message}
}

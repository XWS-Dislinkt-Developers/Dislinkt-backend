package domain

import "time"

type UserConnection struct {
	UserId          int   `bson:"user_id"`
	Private         bool  `bson:"private_profile"`
	Connections     []int `bson:"connections"`
	Requests        []int `bson:"requests"`
	WaitingResponse []int `bson:"waiting_response"`
	Blocked         []int `bson:"blocked"`
}

type Notification struct {
	UserId    int       `bson:"user_id"`
	Content   string    `bson:"content"`
	CreatedAt time.Time `bson:"created_at"`
	Seen      bool      `bson:"seen"`
}

package domain

type UserConnection struct {
	UserId      int   `bson:"user_id"`
	Private     bool  `bson:"private_profile"`
	Connections []int `bson:"connections"`
	Requests    []int `bson:"requests"`
}

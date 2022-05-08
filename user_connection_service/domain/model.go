package domain

type UserConnection struct {
	UserId      int   `bson:"user_id"`
	Connections []int `bson:"connections"`
}

package domain

type UserData struct {
	Id     int    `bson:"id"`
	UserId int    `bson:"user_id"`
	Token  string `bson:"token"`
}

type JobOffer struct {
	Id          int    `bson:"id"`
	UserId      int    `bson:"user_id"`
	Company     string `bson:"company"`
	Position    string `bson:"position"`
	Description string `bson:"description"`
}

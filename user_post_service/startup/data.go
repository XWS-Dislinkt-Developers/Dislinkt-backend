package startup

import (
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_post_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var userPosts = []*domain.UserPost{
	{
		Id:        GetObjectId("623b0cc336a1d6fd8c1cf0f6"),
		UserId:    1,
		CreatedAt: time.Now(),
		Text:      "First post!",
		ImagePath: "imagepath..",

		Reactions: []domain.Reaction{
			{
				UserId:   2,
				Liked:    true,
				Disliked: false,
			},
			{
				UserId:   3,
				Liked:    false,
				Disliked: true,
			},
			{
				UserId:   4,
				Liked:    true,
				Disliked: false,
			},
		},
		Comments: []domain.Comment{
			{
				UserId:    3,
				CreatedAt: time.Now(),
				Text:      "First comment!",
			},
			{
				UserId:    3,
				CreatedAt: time.Now(),
				Text:      "Me again!",
			},
			{
				UserId:    1,
				CreatedAt: time.Now(),
				Text:      "I can comment on my own post!!",
			},
		},
	},
}

func GetObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}

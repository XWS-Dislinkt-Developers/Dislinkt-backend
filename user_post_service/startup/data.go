package startup

import (
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_post_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var userPosts = []*domain.UserPost{
	{
		Id:        GetObjectId("507f1f77bcf86cd799439011"),
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
	{
		Id:        GetObjectId("507f1f77bcf86cd799439012"),
		UserId:    1,
		CreatedAt: time.Now(),
		Text:      "Second post!",
		ImagePath: "imagepath..",

		Reactions: []domain.Reaction{
			{
				UserId:   1,
				Liked:    true,
				Disliked: false,
			},
		},
		Comments: []domain.Comment{
			/*{
				UserId:    3,
				CreatedAt: time.Now(),
				Text:      "Cao!",
			},*/
		},
	},
	{
		Id:        GetObjectId("6276eb70d31c8f2272d2fbe5"),
		UserId:    2,
		CreatedAt: time.Now(),
		Text:      "New post!",
		ImagePath: "imagepath..",

		Reactions: []domain.Reaction{
			{
				UserId:   1,
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
		println(err)
		return objectId
	}
	return primitive.NewObjectID()
}

package startup

import (
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_post_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var userPosts = []*domain.UserPost{
	// UserPost - 01
	{
		Id:        GetObjectId("507f1f77bcf86cd799439011"),
		UserId:    1,
		CreatedAt: time.Now().Add(time.Duration(-25) * time.Hour), // Now - 25h,
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
				UserId:    2,
				CreatedAt: time.Now().Add(time.Duration(-3) * time.Minute), // Now - 3min,
				Text:      "First comment!",
			},
			{
				UserId:    2,
				CreatedAt: time.Now().Add(time.Duration(-2) * time.Minute), // Now - 2min,
				Text:      "Me again!",
			},
			{
				UserId:    1,
				CreatedAt: time.Now().Add(time.Duration(-1) * time.Minute), // Now - 1min,
				Text:      "I can comment on my own post!!",
			},
		},
	},
	// UserPost - 02
	{
		Id:        GetObjectId("507f1f77bcf86cd799439012"),
		UserId:    2,
		CreatedAt: time.Now().Add(time.Duration(-5) * time.Hour), // Now - 5h,
		Text:      "Second post!",
		ImagePath: "imagepath..",

		Reactions: []domain.Reaction{
			{
				UserId:   1,
				Liked:    true,
				Disliked: false,
			},
		},
		Comments: []domain.Comment{},
	},
	// UserPost - 03
	{
		Id:        GetObjectId("6276eb70d31c8f2272d2fbe5"),
		UserId:    4,
		CreatedAt: time.Now().Add(time.Duration(-4) * time.Hour), // Now - 4h,
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
				CreatedAt: time.Now().Add(time.Duration(-50) * time.Minute), // Now - 50sec,
				Text:      "First comment!",
			},
			{
				UserId:    3,
				CreatedAt: time.Now().Add(time.Duration(-40) * time.Minute), // Now - 40sec,
				Text:      "Me again!",
			},
			{
				UserId:    1,
				CreatedAt: time.Now().Add(time.Duration(-20) * time.Minute), // Now - 20sec,
				Text:      "I can comment on my own post!!",
			},
		},
	},
	// UserPost - 04
	{
		Id:        GetObjectId("6276eb70d31c8f2272d2fbe9"),
		UserId:    1,
		CreatedAt: time.Now().Add(time.Duration(-25) * time.Minute), // Now - 25min,
		Text:      "Novi postic!",
		ImagePath: "imagepath..",

		Reactions: []domain.Reaction{
			{
				UserId:   5,
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
				CreatedAt: time.Now().Add(time.Duration(-15) * time.Minute), // Now - 15min,
				Text:      "First comment!",
			},
			{
				UserId:    3,
				CreatedAt: time.Now().Add(time.Duration(-10) * time.Minute), // Now - 10min,
				Text:      "Me again!",
			},
			{
				UserId:    1,
				CreatedAt: time.Now().Add(time.Duration(-8) * time.Minute), // Now - 8min,
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

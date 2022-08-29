package startup

import (
	"time"

	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_post_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var userPosts = []*domain.UserPost{
	// UserPost - 01
	{
		Id:        GetObjectId("507f1f77bcf86cd799439011"),
		UserId:    1,
		CreatedAt: time.Now().Add(time.Duration(-25) * time.Hour), // Now - 25h,
		Text:      "First post!",
		ImagePath: "",
		Likes:     []int{1, 2, 4},
		Dislikes:  []int{},
		Comments: []domain.Comment{
			{
				UserId:    2,
				CreatedAt: time.Now().Add(time.Duration(-3) * time.Minute), // Now - 3min,
				Text:      "First comment!",
			},
			{
				UserId:    2,
				CreatedAt: time.Now().Add(time.Duration(-2) * time.Minute), // Now - 2min,
				Text:      "Me again! With a link now: https://youtu.be/dQw4w9WgXcQ",
			},
			{
				UserId:    1,
				CreatedAt: time.Now().Add(time.Duration(-1) * time.Minute), // Now - 1min,
				Text:      "I can comment on my own post!! 😋",
			},
		},
	},
	// UserPost - 02
	{
		Id:        GetObjectId("507f1f77bcf86cd799439012"),
		UserId:    2,
		CreatedAt: time.Now().Add(time.Duration(-5) * time.Hour), // Now - 5h,
		Text:      "FTN repozitorijum - link: http://www.acs.uns.ac.rs/sr/repozitorijum",
		ImagePath: "",
		Likes:     []int{1},
		Dislikes:  []int{},
		Comments:  []domain.Comment{},
	},
	// UserPost - 03
	{
		Id:        GetObjectId("6276eb70d31c8f2272d2fbe5"),
		UserId:    4,
		CreatedAt: time.Now().Add(time.Duration(-4) * time.Hour), // Now - 4h,
		Text:      "New post!",
		ImagePath: "",
		Likes:     []int{1, 4},
		Dislikes:  []int{3},
		Comments: []domain.Comment{
			{
				UserId:    3,
				CreatedAt: time.Now().Add(time.Duration(-50) * time.Minute), // Now - 50sec,
				Text: "Don't forget to move your body! \n " +
					"8min abs: https://www.youtube.com/watch?v=sWjTnBmCHTY \n" +
					"8min legs: https://youtu.be/B0lF0gDzaAc \n" +
					"8min arms: https://youtu.be/CxYT5cS4ljA \n" +
					"8min buns: https://youtu.be/n538bSON2kA \n" +
					"8min stretch: https://youtu.be/MzVdO8DJyao ",
			},
			{
				UserId:    3,
				CreatedAt: time.Now().Add(time.Duration(-40) * time.Minute), // Now - 40sec,
				Text:      "⠀           (\\ __ /)\n              (UwU)\n       ＿ノ ヽ ノ＼＿ \n  /　`/ ⌒Ｙ⌒ Ｙ　 \\\n( 　(三ヽ人　 /　　|\n|　ﾉ⌒＼ ￣￣ヽ　 ノ\nヽ＿＿＿＞､＿＿／\n         ｜( 王 ﾉ〈 \n         /ﾐ`ー―彡\\ \n       / ╰           ╯  \\",
			},
			{
				UserId:    1,
				CreatedAt: time.Now().Add(time.Duration(-20) * time.Minute), // Now - 20sec,
				Text:      "YEEEAA!! 💪💪💪💪💪",
			},
		},
	},
	// UserPost - 04
	{
		Id:        GetObjectId("6276eb70d31c8f2272d2fbe9"),
		UserId:    1,
		CreatedAt: time.Now().Add(time.Duration(-25) * time.Minute), // Now - 25min,
		Text:      "Johnny Depp won! Objection! Hearsay! Here is a parody: https://youtu.be/mDT91ih6S5Q",
		ImagePath: "",
		Likes:     []int{5, 4},
		Dislikes:  []int{3},
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

package api

import (
	pb_post "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/user_post_service"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_post_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// UserPostMapper - Za preuzimanje postojecih postova (zato ucitavamo nizove: Reactions, Comments)
func mapUserPost(userPost *domain.UserPost) *pb_post.UserPost {
	userPostPb := &pb_post.UserPost{
		Id:        userPost.Id.Hex(),
		UserId:    int64(userPost.UserId),
		CreatedAt: timestamppb.New(userPost.CreatedAt),
		Text:      userPost.Text,
		ImagePath: userPost.ImagePath,
		Likes:     make([]int64, 0),
		Dislikes:  make([]int64, 0),
		Comments:  make([]*pb_post.Comment, 0),
	}
	for _, like := range userPost.Likes {
		userPostPb.Likes = append(userPostPb.Likes, int64(like))
	}
	for _, dislike := range userPost.Dislikes {
		userPostPb.Dislikes = append(userPostPb.Dislikes, int64(dislike))
	}
	for _, comment := range userPost.Comments {
		commentPb := &pb_post.Comment{
			UserId:    int64(comment.UserId),
			CreatedAt: timestamppb.New(comment.CreatedAt),
			Text:      comment.Text,
		}
		userPostPb.Comments = append(userPostPb.Comments, commentPb)
	}
	return userPostPb
}

// UserPostMapper - Kad kreiramo userPost (zato su prazni nizovi: Reactions, Comments)
func mapNewUserPost(userPostPb *pb_post.UserPost, userId int) *domain.UserPost {
	userPost := &domain.UserPost{
		Id:        GetObjectId(userPostPb.Id),
		UserId:    userId,
		Text:      userPostPb.Text,
		ImagePath: userPostPb.ImagePath,
		CreatedAt: timestamppb.Now().AsTime(), // Ovako cuvamo trenutno vreme
		Likes:     make([]int, 0),             // Kreira prazan niz
		Dislikes:  make([]int, 0),             // Kreira prazan niz
		Comments:  make([]domain.Comment, 0),  // Kreira prazan niz
	}
	return userPost
}

// Lale: Mislim da potrebno napraviti mapNewReactionToUserPost, mapNewCommentToUserPost
//TODO: mapNewReactionToUserPost(), mapNewCommentToUserPost()
func mapNewCommentToUserPost(userCommentPb *pb_post.AddCommentRequest, userId int) *domain.Comment {
	comment := &domain.Comment{
		UserId:    userId,
		CreatedAt: timestamppb.Now().AsTime(),
		Text:      userCommentPb.AddComment.Text,
	}
	return comment
}

func converterInt64ToIntArray(likes []int64) []int {
	ret := make([]int, 0)
	for i := range likes {
		ret[i] = int(likes[i])
	}
	return ret
}

func GetObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}

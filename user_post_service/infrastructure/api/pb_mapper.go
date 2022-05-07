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
		Reactions: make([]*pb_post.Reaction, 0),
		Comments:  make([]*pb_post.Comment, 0),
	}
	for _, reaction := range userPost.Reactions {
		reactionPb := &pb_post.Reaction{
			UserId:   int64(reaction.UserId),
			Liked:    reaction.Liked,
			Disliked: reaction.Disliked,
		}
		userPostPb.Reactions = append(userPostPb.Reactions, reactionPb)
	}
	for _, comment := range userPost.Comments {
		commentPb := &pb_post.Comment{
			UserId:    int64(comment.UserId),
			CreatedAt: timestamppb.New(userPost.CreatedAt),
			Text:      comment.Text,
		}
		userPostPb.Comments = append(userPostPb.Comments, commentPb)
	}
	return userPostPb
}

// UserPostMapper - Kad kreiramo userPost (zato su prazni nizovi: Reactions, Comments)
func mapNewUserPost(userPostPb *pb_post.UserPost, userId int) *domain.UserPost {
	userPost := &domain.UserPost{
		Id:        GetObjectId(userPostPb.Id), // Ako sam dobro razumeo, generisace novi Id.
		UserId:    userId,                     /*int(userPostPb.UserId)*/
		Text:      userPostPb.Text,
		ImagePath: userPostPb.ImagePath,
		CreatedAt: timestamppb.Now().AsTime(), // Ovako cuvamo trenutno vreme
		Reactions: make([]domain.Reaction, 0), // Kreira prazan niz
		Comments:  make([]domain.Comment, 0),  // Kreira prazan niz
	}
	/*
		-- Mislim da ovo ne treba
			-- Kad se kreira user-post, nemamo comments ni reactions.
		for _, reactionPb := range userPostPb.Reactions {
			reaction := domain.Reaction{
				UserId:   int(reactionPb.UserId),
				IsItLike: reactionPb.IsItLike,
			}
			userPost.Reactions = append(userPost.Reactions, reaction)
		}
		for _, commentPb := range userPostPb.Comments {
			comment := domain.Comment{
				UserId:    int(commentPb.UserId),
				CreatedAt: commentPb.CreatedAt.AsTime(),
				Text:      commentPb.Text,
			}
			userPost.Comments = append(userPost.Comments, comment)
		}
	*/
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

func GetObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}

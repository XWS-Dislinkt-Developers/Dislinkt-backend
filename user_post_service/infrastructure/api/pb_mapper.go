package api

import (
	pb "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/user_post_service"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_post_service/domain"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapUserPost(userPost *domain.UserPost) *pb.UserPost {
	userPostPb := &pb.UserPost{
		Id:        userPost.Id,     // TODO: HEX??? .Hex()
		UserId:    userPost.UserId, // TODO: HEX???
		CreatedAt: timestamppb.New(userPost.CreatedAt),
		Text:      userPost.Text,
		ImagePath: userPost.ImagePath,
		Reactions: make([]*pb.Reaction, 0),
		Comments:  make([]*pb.Comment, 0),
	}
	for _, reaction := range userPost.Reactions {
		reactionPb := &pb.Reaction{
			UserId:   reaction.UserId.Hex(), // TODO: HEX ??
			IsItLike: reaction.IsItLike,
		}
		userPostPb.Reactions = append(userPostPb.Reactions, reactionPb)
	}
	for _, comment := range userPost.Comments {
		commentPb := &pb.Comment{
			UserId:    comment.UserId, // TODO: HEX ??
			CreatedAt: timestamppb.New(userPost.CreatedAt),
			Text:      comment.Text,
		}
		userPostPb.Comments = append(userPostPb.Comments, commentPb)
	}
	return userPostPb
}

func mapNewUserPost(userPostPb *pb.UserPost) *domain.UserPost {
	userPost := &domain.UserPost{
		// TODO: CreatedAt: Time.Now??
		Reactions: make([]domain.Reaction, 0),
		Comments:  make([]domain.Comment, 0),
	}
	for _, reactionPb := range userPostPb.Reactions {
		reaction := domain.Reaction{
			UserId:   reactionPb.UserId,
			IsItLike: reactionPb.IsItLike,
		}
		userPost.Reactions = append(userPost.Reactions, reaction)
	}
	for _, commentPb := range userPostPb.Comments {
		comment := domain.Comment{
			UserId:    commentPb.UserId,
			CreatedAt: timestamppb.New(comment.CreatedAt),
			Text:      comment.Text,
		}
		userPost.Comments = append(userPost.Comments, comment)
	}
	return userPost
}

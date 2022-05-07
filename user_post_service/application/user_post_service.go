package application

import (
	"fmt"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_post_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserPostService struct {
	store domain.UserPostStore
}

func NewUserPostService(store domain.UserPostStore) *UserPostService {
	return &UserPostService{
		store: store,
	}
}

func (service *UserPostService) Get(id primitive.ObjectID) (*domain.UserPost, error) {
	return service.store.Get(id)
}

func (service *UserPostService) GetAll() ([]*domain.UserPost, error) {
	return service.store.GetAll()
}

func (service *UserPostService) Create(userPost *domain.UserPost) error {
	userPost.CreatedAt = time.Now()
	err := service.store.Insert(userPost)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserPostService) AddComment(comment *domain.Comment, idPost primitive.ObjectID) (*domain.UserPost, error) {
	UserPost, _ := service.store.Get(idPost)
	UserPost.Comments = append(UserPost.Comments, *comment)
	fmt.Printf("Brojim koliko ima komentara ", len(UserPost.Comments))
	fmt.Printf("U user service sam, citam userpostid ", UserPost.Id.Hex())
	service.store.UpdateComments(UserPost)
	return service.store.Get(idPost)
}

func (service *UserPostService) AddReaction(reaction *domain.Reaction, idPost primitive.ObjectID) (*domain.UserPost, error) {
	UserPost, _ := service.store.Get(idPost)
	userAlreadyReacted := false
	for _, r := range UserPost.Reactions {
		if r.UserId == reaction.UserId {
			userAlreadyReacted = true
			println(userAlreadyReacted)

		}
	}
	if userAlreadyReacted {
		println("NASAO SAM USEERA KOJI JE REAGOVAO")
		service.UpdateReaction(reaction, UserPost)
	} else {
		UserPost.Reactions = append(UserPost.Reactions, *reaction)
		service.store.AddReaction(UserPost)
	}

	println(UserPost)
	return service.store.Get(idPost)
}

func (service *UserPostService) UpdateReaction(reaction *domain.Reaction, userPost *domain.UserPost) {
	var updatedReaction *domain.Reaction
	for _, r := range userPost.Reactions {
		if r.UserId == reaction.UserId {
			if r.Liked && !r.Disliked && reaction.Liked && !reaction.Disliked {
				r.Liked = false
				updatedReaction = &r
				break
			}
			if !r.Liked && r.Disliked && !reaction.Liked && reaction.Disliked {
				r.Disliked = false
				updatedReaction = &r
				break
			}
			if !r.Liked && !r.Disliked && reaction.Liked && !reaction.Disliked {
				r.Liked = true
				updatedReaction = &r
				break
			}
			if !r.Liked && !r.Disliked && !reaction.Liked && reaction.Disliked {
				r.Disliked = true
				updatedReaction = &r
				break
			}
			if !reaction.Liked && !reaction.Disliked {
				updatedReaction = &r
				break
			}
			if reaction.Liked && reaction.Disliked {
				updatedReaction = &r
				break
			}
		}

	}
	service.store.UpdateReactions(updatedReaction, userPost)
}

package application

import (
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_post_service/domain"
	logg "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_post_service/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
	"time"
)

type UserPostService struct {
	store       domain.UserPostStore
	loggerInfo  *logg.Logger
	loggerError *logg.Logger
}

func NewUserPostService(store domain.UserPostStore, loggerInfo *logg.Logger, loggerError *logg.Logger) *UserPostService {
	return &UserPostService{
		store:       store,
		loggerInfo:  loggerInfo,
		loggerError: loggerError,
	}
}

func (service *UserPostService) Get(id primitive.ObjectID) (*domain.UserPost, error) {
	return service.store.Get(id)
}
func (service *UserPostService) GetAll() ([]*domain.UserPost, error) {
	return service.store.GetAll()
}
func (service *UserPostService) GetUserPosts(idUser int) ([]*domain.UserPost, error) {
	return service.store.GetPostsByUserId(idUser)
}
func (service *UserPostService) GetPostsForLoggedUserProfile(idUser int) ([]*domain.UserPost, error) {
	posts := make([]*domain.UserPost, 0)
	allPosts, _ := service.GetAll()

	for _, post := range allPosts {
		if post.UserId == idUser {
			posts = append(posts, post)
		}
	}

	return posts, nil
}

func (service *UserPostService) Create(userPost *domain.UserPost) error {
	userPost.CreatedAt = time.Now()
	err := service.store.Insert(userPost)
	service.loggerInfo.Logger.Infof("User_post_service: USCNP | UI  " + strconv.Itoa(userPost.UserId))
	if err != nil {
		service.loggerError.Logger.Errorf("User_post_service: UFCNP | UI  " + strconv.Itoa(userPost.UserId))
		return err
	}
	return nil
}

func (service *UserPostService) AddComment(comment *domain.Comment, idPost primitive.ObjectID) (*domain.UserPost, error) {
	UserPost, _ := service.store.Get(idPost)
	if comment.Text != "" {
		// comment.CreatedAt = time.Now()
		UserPost.Comments = append(UserPost.Comments, *comment)
		service.loggerInfo.Logger.Infof("User_post_service: USANCTP | UI " + strconv.Itoa(UserPost.UserId))
		service.store.UpdateComments(UserPost)
	}
	return service.store.Get(idPost)
}
func (service *UserPostService) Like(idLoggedUser int, idPost primitive.ObjectID) (*domain.UserPost, error) {
	userPost, _ := service.store.Get(idPost)
	if service.likedByUser(idLoggedUser, userPost) {
		// If it's already liked, the post won't be liked neither disliked
		userPost.Likes = findAndDelete(userPost.Likes, idLoggedUser)
		service.store.UpdateLikes(userPost)
	} else if service.dislikedByUser(idLoggedUser, userPost) {
		// If it's already disliked, the post will be liked and won't be disliked
		userPost.Dislikes = findAndDelete(userPost.Dislikes, idLoggedUser)
		userPost.Likes = append(userPost.Likes, idLoggedUser)
		service.store.UpdateLikes(userPost)
		service.store.UpdateDislikes(userPost)
	} else {
		// If the post will be liked
		userPost.Likes = append(userPost.Likes, idLoggedUser)
		service.store.UpdateLikes(userPost)
	}
	return service.store.Get(idPost)
}
func (service *UserPostService) Dislike(idLoggedUser int, idPost primitive.ObjectID) (*domain.UserPost, error) {
	userPost, _ := service.store.Get(idPost)
	if service.dislikedByUser(idLoggedUser, userPost) {
		// If it's already disliked, the post won't be liked neither disliked
		userPost.Dislikes = findAndDelete(userPost.Likes, idLoggedUser)
		service.store.UpdateDislikes(userPost)
	} else if service.likedByUser(idLoggedUser, userPost) {
		// If it's already liked, the post will be disliked and won't be liked
		userPost.Likes = findAndDelete(userPost.Likes, idLoggedUser)
		userPost.Dislikes = append(userPost.Dislikes, idLoggedUser)
		service.store.UpdateLikes(userPost)
		service.store.UpdateDislikes(userPost)
	} else {
		// If the post will be liked
		userPost.Dislikes = append(userPost.Dislikes, idLoggedUser)
		service.store.UpdateDislikes(userPost)
	}
	return service.store.Get(idPost)
}

func (service *UserPostService) likedByUser(idLoggedUser int, userPost *domain.UserPost) bool {
	for _, l := range userPost.Likes {
		if l == idLoggedUser {
			return true
		}
	}
	return false
}
func (service *UserPostService) dislikedByUser(idLoggedUser int, userPost *domain.UserPost) bool {
	for _, l := range userPost.Dislikes {
		if l == idLoggedUser {
			return true
		}
	}
	return false
}

func findAndDelete(s []int, item int) []int {
	index := 0
	for _, i := range s {
		if i != item {
			s[index] = i
			index++
		}
	}
	return s[:index]
}

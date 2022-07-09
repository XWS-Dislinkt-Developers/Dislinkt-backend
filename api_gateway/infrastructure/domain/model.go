package domain

import (
	"time"
)

type User struct {
	ID              int    `json:"id" ,gorm:"primaryKey,autoIncrement:true" `
	UserId          int    `json:"userId"  `
	Name            string `json:"name" validate:"required,name"`
	Username        string `json:"username" ,gorm:"unique" validate:"required,username" `
	Password        string `json:"password"`
	Email           string `json:"email" ,gorm:"unique" validate:"required,email"`
	ConfirmPassword string `json:"confirmPassword"`
	Gender          string `json:"gender"`
}

type UserDTO struct {
	ID               int
	UserId           int
	Name             string
	Username         string
	Gender           string
	IsPrivateProfile bool
	Biography        string
}

type PasswordRecoveryDTO struct {
	Code            string
	Password        string
	ConfirmPassword string
}

type Reaction struct {
	UserId   int
	Liked    bool
	Disliked bool
}

type Comment struct {
	UserId    int
	CreatedAt time.Time
	Text      string
}

type UserPost struct {
	Id        string
	UserId    int
	CreatedAt time.Time
	Text      string
	ImagePath string
	Reactions []Reaction
	Comments  []Comment
}

type Fs struct {
	Feed []UserPost
}

type Users struct {
	Users []UserDTO
}

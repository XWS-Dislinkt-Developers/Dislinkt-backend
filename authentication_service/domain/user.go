package domain

import "time"

type User struct {
	ID               int       `json:"id" ,gorm:"primaryKey,;autoIncrement:true" `
	Name             string    `json:"name" validate:"required,name"`
	Username         string    `json:"username" ,gorm:"unique" validate:"required,username" `
	Password         string    `json:"password"`
	Email            string    `json:"email" ,gorm:"unique" validate:"required,email"`
	Address          string    `json:"address"`
	Gender           string    `json:"gender"`
	DateOfBirth      time.Time `json:"dateOfBirth"`
	Biography        string    `json:"biography"`
	PhoneNumber      string    `json:"phoneNumber"`
	IsPrivateProfile bool      `json:"isPrivateProfile"`
	Work             string    `json:"work"`
	Education        string    `json:"education"`
	Interests        string    `json:"interests"`
	Skills           string    `json:"skills"`
}

type UpdateUserDto struct {
	Username    string    `json:"username" ,gorm:"unique" validate:"required,username" `
	Name        string    `json:"name" validate:"required,name"`
	PhoneNumber string    `json:"phoneNumber"`
	Email       string    `json:"email" ,gorm:"unique" validate:"required,email"`
	Gender      string    `json:"gender"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Biography   string    `json:"BateOfBirth"`
}

type UpdateUserSAIDto struct {
	Skills    string `json:"skills"`
	Interests string `json:"interests"`
}

type UpdateUserWAEDto struct {
	Work      string `json:"work"`
	Education string `json:"education"`
}

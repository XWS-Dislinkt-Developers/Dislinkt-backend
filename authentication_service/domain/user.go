package domain

type User struct {
	ID            int    `json:"id" ,gorm:"primaryKey,autoIncrement:true" `
	Username      string `json:"username" ,gorm:"unique" validate:"required,username" `
	Password      string `json:"password"`
	Email         string `json:"email" ,gorm:"unique" validate:"required,email"`
	IsItConfirmed bool   `json:"isItConfirmed"`
	Role          string `json:"role""`
}

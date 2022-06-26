package domain

type User struct {
	ID              int    `json:"id" ,gorm:"primaryKey,autoIncrement:true" `
	Name            string `json:"name" validate:"required,name"`
	Username        string `json:"username" ,gorm:"unique" validate:"required,username" `
	Password        string `json:"password"`
	Email           string `json:"email" ,gorm:"unique" validate:"required,email"`
	ConfirmPassword string `json:"confirmPassword"`
	Gender          string `json:"gender"`
}

type PasswordRecoveryDTO struct {
	Code            string
	Password        string
	ConfirmPassword string
}

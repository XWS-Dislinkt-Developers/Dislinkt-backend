package domain

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name" validate:"required,name"`
	Username string `json:"username" ,gorm:"unique" validate:"required,username" `
	Password string `json:"password"`
	Email    string `json:"email" ,gorm:"unique" validate:"required,email"`
	Address  string `json:"address"`
	Gender   string `json:"gender"`
}

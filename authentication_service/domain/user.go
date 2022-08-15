package domain

type User struct {
	ID            int                `json:"id" ,gorm:"primaryKey,autoIncrement:true" `
	Username      string             `json:"username" ,gorm:"unique" validate:"required,username" `
	Password      string             `json:"password"`
	Email         string             `json:"email" ,gorm:"unique" validate:"required,email"`
	IsItConfirmed bool               `json:"isItConfirmed"`
	Role          string             `json:"role""`
	Status        RegistrationStatus `json:"status"`
}
type RegistrationStatus int8

const (
	Pending RegistrationStatus = iota
	Approved
	Cancelled
)

func (status RegistrationStatus) String() string {
	switch status {
	case Pending:
		return "Pending"
	case Approved:
		return "Approved"
	case Cancelled:
		return "Cancelled"
	}
	return "Unknown"
}

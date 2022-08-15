package register_user

type UserDetails struct {
	ID            int64
	Name          string
	Username      string
	Password      string
	Email         string
	IsItConfirmed bool
	Role          string
}

type RegisterUserCommandType int8

const (
	SaveInAuthDB RegisterUserCommandType = iota
	RollbackSavingAuth
	CancelRegistration
	SendEmail
	SaveInUserDB
	ApproveRegistration
	UnknownCommand
)

type RegisterUserCommand struct {
	User UserDetails
	Type RegisterUserCommandType
}

type RegisterUserReplyType int8

const (
	UserSavedDB RegisterUserReplyType = iota
	UserNotSavedDB

	UserNotSavedAuth
	UserSavedAuth
	UserAuthRollback
	UserRegistrationCancelled
	EmailFailed
	EmailSent
	UserRegistered
	UnknownReply
)

type RegisterUserReply struct {
	User UserDetails
	Type RegisterUserReplyType
}

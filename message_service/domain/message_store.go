package domain

type MessageStore interface {
	// CRUD - READ method(s)
	GetAll() ([]*Message, error)
	GetAllUsersMessagesByUserId(id int) ([]*Message, error)
	/*
		GetAllReceiversMessagesByUserId(id int) ([]*Message, error)
		GetAllMessagesBetweenUsers(userId1, userId2 int) ([]*Message, error)
	*/
	// CRUD - CREATE method(s)
	Insert(message *Message) error
	// CRUD - DELETE method(s)
	DeleteAll()
}

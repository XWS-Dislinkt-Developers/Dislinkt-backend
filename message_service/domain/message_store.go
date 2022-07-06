package domain

type MessageStore interface {
	GetBySenderId(id int) (*Message, error)
	GetByReceiverId(id int) (*Message, error)
	GetAll() ([]*Message, error)
	Insert(message *Message) error
}

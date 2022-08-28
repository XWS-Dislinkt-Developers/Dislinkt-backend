package nats

import (
	saga "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/saga/messaging"
	"github.com/nats-io/nats.go"
)

type Subscriber struct {
	conn       *nats.EncodedConn
	subject    string
	queueGroup string
}

func NewNATSSubscriber(host, port, user, password, subject, queueGroup string) (saga.Subscriber, error) {
	conn, err := getConnection(host, port, user, password)
	if err != nil {
		return nil, err
		println("Greska kod subscribera za user service")
		println(err.Error())
		println(err.Error())
		println(err.Error())
	}
	encConn, err := nats.NewEncodedConn(conn, nats.JSON_ENCODER)
	if err != nil {
		return nil, err
	}
	return &Subscriber{
		conn:       encConn,
		subject:    subject,
		queueGroup: queueGroup,
	}, nil
}

func (s *Subscriber) Subscribe(handler interface{}) error {
	_, err := s.conn.QueueSubscribe(s.subject, s.queueGroup, handler)
	if err != nil {
		return err
	}
	return nil
}

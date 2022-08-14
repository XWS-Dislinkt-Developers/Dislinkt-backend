package api

import (
	pb_message "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/message_service"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/message_service/domain"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapMessage(message *domain.Message) *pb_message.Message {
	messagePb := &pb_message.Message{
		SenderId:   int64(message.SenderId),
		ReceiverId: int64(message.ReceiverId),
		Content:    message.Content,
		CreatedAt:  timestamppb.New(message.CreatedAt),
	}
	return messagePb
}
func mapNewMessage(messagePb *pb_message.Message) *domain.Message {
	message := &domain.Message{
		SenderId:   int(messagePb.SenderId),
		ReceiverId: int(messagePb.ReceiverId),
		CreatedAt:  timestamppb.Now().AsTime(),
		Content:    messagePb.Content,
	}
	return message
}

package startup

import (
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/message_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var messages = []*domain.Message{
	//-----[000 -> 001]--[PERA -> PERA]--
	{
		SenderId:   1,
		ReceiverId: 1,
		Content:    "I'm talking to myself!",
		CreatedAt:  time.Now().Add(time.Duration(-52) * time.Minute), // Now - 52min
	},
	{
		SenderId:   1,
		ReceiverId: 1,
		Content:    "I'm talking to myself AGAIN!",
		CreatedAt:  time.Now().Add(time.Duration(-52) * time.Minute), // Now - 52min
	},
	//-----[002 -> 005]--[PERA -> JOKA]--
	{
		SenderId:   1,
		ReceiverId: 2,
		Content:    "Hi Joka! I'm Pera! What's up?",
		CreatedAt:  time.Now().Add(time.Duration(-50) * time.Minute), // Now - 50min
	},
	{
		SenderId:   2,
		ReceiverId: 1,
		Content:    "Hi Pera! Nice to meet you!",
		CreatedAt:  time.Now().Add(time.Duration(-49) * time.Minute), // Now - 49min
	},
	{
		SenderId:   2,
		ReceiverId: 1,
		Content:    "Everything is awesome!",
		CreatedAt:  time.Now().Add(time.Duration(-48) * time.Minute), // Now - 48min
	},
	{
		SenderId:   1,
		ReceiverId: 2,
		Content:    "That's good to hear! :))",
		CreatedAt:  time.Now().Add(time.Duration(-47) * time.Minute), // Now - 47min
	},
	//-----[006 -> 009]--[PERA -> MARKO]--
	{
		SenderId:   1,
		ReceiverId: 3,
		Content:    "Hi Marko! I'm Pera! What's up?",
		CreatedAt:  time.Now().Add(time.Duration(-46) * time.Minute), // Now - 46min
	},
	{
		SenderId:   3,
		ReceiverId: 1,
		Content:    "Hi Pera! I'm Marko! Nice to meet you!",
		CreatedAt:  time.Now().Add(time.Duration(-45) * time.Minute), // Now - 45min
	},
	{
		SenderId:   3,
		ReceiverId: 1,
		Content:    "Everything is awesome!",
		CreatedAt:  time.Now().Add(time.Duration(-44) * time.Minute), // Now - 44min
	},
	{
		SenderId:   1,
		ReceiverId: 3,
		Content:    "That's good to hear! :))",
		CreatedAt:  time.Now().Add(time.Duration(-43) * time.Minute), // Now - 43min
	},
	//-----[010 -> 013]--[PERA -> ZEKSA]--
	{
		SenderId:   1,
		ReceiverId: 4,
		Content:    "Hi Zeksa! I'm Pera! What's up?",
		CreatedAt:  time.Now().Add(time.Duration(-42) * time.Minute), // Now - 42min
	},
	{
		SenderId:   4,
		ReceiverId: 1,
		Content:    "Hi Pera! I'm Zeksa! Nice to meet you!",
		CreatedAt:  time.Now().Add(time.Duration(-41) * time.Minute), // Now - 41min
	},
	{
		SenderId:   4,
		ReceiverId: 1,
		Content:    "Everything is awesome!",
		CreatedAt:  time.Now().Add(time.Duration(-40) * time.Minute), // Now - 40min
	},
	{
		SenderId:   1,
		ReceiverId: 4,
		Content:    "That's good to hear! :))",
		CreatedAt:  time.Now().Add(time.Duration(-39) * time.Minute), // Now - 39min
	},
	//-----[014 -> 017]--[PERA -> SANJA]--
	{
		SenderId:   1,
		ReceiverId: 5,
		Content:    "Hi Sanja! I'm Pera! What's up?",
		CreatedAt:  time.Now().Add(time.Duration(-38) * time.Minute), // Now - 38min
	},
	{
		SenderId:   5,
		ReceiverId: 1,
		Content:    "Hi Pera! I'm Sanja! Nice to meet you!",
		CreatedAt:  time.Now().Add(time.Duration(-37) * time.Minute), // Now - 37min
	},
	{
		SenderId:   5,
		ReceiverId: 1,
		Content:    "Everything is awesome!",
		CreatedAt:  time.Now().Add(time.Duration(-36) * time.Minute), // Now - 36min
	},
	{
		SenderId:   1,
		ReceiverId: 5,
		Content:    "That's good to hear! :))",
		CreatedAt:  time.Now().Add(time.Duration(-35) * time.Minute), // Now - 35min
	},
	//-----[018 -> 021]--[PERA -> TANJA]--
	{
		SenderId:   1,
		ReceiverId: 6,
		Content:    "Hi Tanja! I'm Pera! What's up?",
		CreatedAt:  time.Now().Add(time.Duration(-34) * time.Minute), // Now - 34min
	},
	{
		SenderId:   6,
		ReceiverId: 1,
		Content:    "Hi Pera! I'm Tanja! Nice to meet you!",
		CreatedAt:  time.Now().Add(time.Duration(-33) * time.Minute), // Now - 33min
	},
	{
		SenderId:   6,
		ReceiverId: 1,
		Content:    "Everything is awesome!",
		CreatedAt:  time.Now().Add(time.Duration(-32) * time.Minute), // Now - 32min
	},
	{
		SenderId:   1,
		ReceiverId: 6,
		Content:    "That's good to hear! :))",
		CreatedAt:  time.Now().Add(time.Duration(-31) * time.Minute), // Now - 31min
	},
	//-----[022 -> 025]--[PERA -> LALE]--
	{
		SenderId:   1,
		ReceiverId: 7,
		Content:    "Hi Lale! I'm Pera! What's up?",
		CreatedAt:  time.Now().Add(time.Duration(-30) * time.Minute), // Now - 30min
	},
	{
		SenderId:   7,
		ReceiverId: 1,
		Content:    "Hi Pera! I'm Lale! Nice to meet you!",
		CreatedAt:  time.Now().Add(time.Duration(-29) * time.Minute), // Now - 29min
	},
	{
		SenderId:   7,
		ReceiverId: 1,
		Content:    "Everything is awesome!",
		CreatedAt:  time.Now().Add(time.Duration(-28) * time.Minute), // Now - 28min
	},
	{
		SenderId:   1,
		ReceiverId: 7,
		Content:    "That's good to hear! :))",
		CreatedAt:  time.Now().Add(time.Duration(-27) * time.Minute), // Now - 27min
	},
	//-----[026 -> 029]--[PERA -> NENA]--
	{
		SenderId:   1,
		ReceiverId: 8,
		Content:    "Hi Nena! I'm Pera! What's up?",
		CreatedAt:  time.Now().Add(time.Duration(-26) * time.Minute), // Now - 26min
	},
	{
		SenderId:   8,
		ReceiverId: 1,
		Content:    "Hi Pera! I'm Nena! Nice to meet you!",
		CreatedAt:  time.Now().Add(time.Duration(-25) * time.Minute), // Now - 25min
	},
	{
		SenderId:   8,
		ReceiverId: 1,
		Content:    "Everything is awesome!",
		CreatedAt:  time.Now().Add(time.Duration(-24) * time.Minute), // Now - 24min
	},
	{
		SenderId:   1,
		ReceiverId: 8,
		Content:    "That's good to hear! :))",
		CreatedAt:  time.Now().Add(time.Duration(-23) * time.Minute), // Now - 23min
	},
}

func GetObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		println(err)
		return objectId
	}
	return primitive.NewObjectID()
}

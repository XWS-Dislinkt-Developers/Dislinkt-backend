package register_user

type Color struct {
	Code string
}

type Product struct {
	Id    string
	Color Color
}

type OrderItem struct {
	Product  Product
	Quantity uint16
}

type OrderDetails struct {
	Id      string
	Items   []OrderItem
	Address string
}

type UserDetails struct {
	Id            int
	Username      string
	Password      string
	Email         string
	IsItConfirmed bool
	Role          string
}

type RegisterUserCommandType int8

type CreateOrderCommandType int8

const (
	RegisterUser RegisterUserCommandType = iota
	RollbackUser
	ApproveUser
	CancelUser
	UnknownCommand
)

type CreateOrderCommand struct {
	Order OrderDetails
	Type  CreateOrderCommandType
}

type RegisterUserCommand struct {
	User UserDetails
	Type RegisterUserCommandType
}

type CreateOrderReplyType int8

type RegisterUserReplyType int8

const (
	UserSaved RegisterUserReplyType = iota
	UserNotSaved
	InventoryRolledBack
	OrderShippingScheduled
	OrderShippingNotScheduled
	OrderApproved
	OrderCancelled
	UnknownReply
)

type RegisterUserReply struct {
	User UserDetails
	Type RegisterUserCommandType
}

type CreateOrderReply struct {
	Order OrderDetails
	Type  CreateOrderCommandType
}

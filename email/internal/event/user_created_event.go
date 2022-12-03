package event

type UserCreatedEvent struct {
	UserId    int    `json:"userId"`
	UserEmail string `json:"userEmail"`
}

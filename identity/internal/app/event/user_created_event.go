package event

type UserCreated struct {
	UserId    int    `json:"userId"`
	UserEmail string `json:"userEmail"`
}

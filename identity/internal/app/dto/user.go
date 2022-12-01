package dto

type CreateUserRequest struct {
	FirstName   string      `json:"firstName"`
	LastName    string      `json:"lastName"`
	Username    string      `json:"username"`
	Email       string      `json:"email"`
	PhoneNumber string      `json:"phoneNumber"`
	Password    string      `json:"password"`
	UserRoles   []*UserRole `json:"userRoles"`
}

type UserRole struct {
	RoleId int `json:"roleId"`
}

type GetUserResponse struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

package usecase

import "github.com/ysfglmzz/e-shop-microservices/identity/internal/app/dto"

type IIdentityService interface {
	CreateUser(createUserRequest dto.CreateUserRequest) error
	LoginUser(loginUserRequest dto.LoginUserRequest) (*dto.TokenResponse, error)
	VerifyUserByCode(code string) error
}

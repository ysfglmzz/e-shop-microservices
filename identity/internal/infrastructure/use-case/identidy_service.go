package usecase

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/ysfglmzz/e-shop-microservices/identity/internal/app/dto"
	"github.com/ysfglmzz/e-shop-microservices/identity/internal/app/event"
	"github.com/ysfglmzz/e-shop-microservices/identity/internal/app/model"
	"github.com/ysfglmzz/e-shop-microservices/identity/internal/app/repository"
	usecase "github.com/ysfglmzz/e-shop-microservices/identity/internal/app/use-case"
	"golang.org/x/crypto/bcrypt"
)

type IdentityService struct {
	messageService     usecase.IMessageService
	idendityRepository repository.IIdentityRepository
}

func NewIdentityService(
	messageService usecase.IMessageService,
	idendityRepository repository.IIdentityRepository) *IdentityService {
	return &IdentityService{messageService: messageService, idendityRepository: idendityRepository}
}

func (i *IdentityService) CreateUser(createUserRequest dto.CreateUserRequest) error {
	var userModel model.User
	if err := copier.Copy(&userModel, createUserRequest); err != nil {
		return err
	}
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(createUserRequest.Password), bcrypt.DefaultCost)
	userModel.PasswordHash = string(hashPassword)
	if err := i.idendityRepository.AddUser(&userModel); err != nil {
		return err
	}
	event := event.UserCreated{UserId: userModel.Id, UserEmail: userModel.Email}
	return i.messageService.PublishUserCreatedEvent(event)
}

func (i *IdentityService) LoginUser(loginUserRequest dto.LoginUserRequest) (*dto.TokenResponse, error) {
	userModel, err := i.idendityRepository.GetUserByEmail(loginUserRequest.Email)
	if err != nil {
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(userModel.PasswordHash), []byte(loginUserRequest.Password)); err != nil {
		return nil, err
	}

	return i.createToken(*userModel)
}

func (i *IdentityService) createToken(user model.User) (*dto.TokenResponse, error) {
	expirationTime := time.Now().Add(time.Minute * time.Duration(600))
	claims := jwt.MapClaims{}
	newUUID := uuid.New()
	roles, err := i.idendityRepository.GetUserRolesByUserId(user.Id)
	if err != nil {
		return nil, err
	}
	claims["uuid"] = newUUID.String()
	claims["exp"] = expirationTime.Unix()
	claims["user_id"] = user.Id
	claims["roles"] = roles

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := newToken.SignedString([]byte("abc"))
	if err != nil {
		return nil, err
	}

	newTokenDetail := model.TokenDetail{
		UUID:           newUUID,
		UserId:         user.Id,
		Token:          tokenString,
		ExpirationDate: expirationTime,
	}

	if err = i.idendityRepository.AddToken(&newTokenDetail); err != nil {
		return nil, err
	}
	return &dto.TokenResponse{Token: tokenString, ExpirationDate: expirationTime}, nil
}

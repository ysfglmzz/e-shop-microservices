package usecase

import (
	"crypto/rand"
	"io"
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

type identityService struct {
	tokenSecretKey      string
	tokenExpirationTime int
	messageService      usecase.IMessageService
	idendityRepository  repository.IIdentityRepository
}

func NewIdentityService(
	messageService usecase.IMessageService,
	idendityRepository repository.IIdentityRepository,
	tokenSecretKey string,
	tokenExpirationTime int,
) *identityService {
	return &identityService{messageService: messageService, idendityRepository: idendityRepository}
}

func (i *identityService) CreateUser(createUserRequest dto.CreateUserRequest) error {
	var userModel model.User
	if err := copier.Copy(&userModel, createUserRequest); err != nil {
		return err
	}
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(createUserRequest.Password), bcrypt.DefaultCost)
	userModel.PasswordHash = string(hashPassword)
	userModel.VerificationCode = generatRandom6digits()
	if err := i.idendityRepository.AddUser(&userModel); err != nil {
		return err
	}
	event := event.UserCreated{UserId: userModel.Id, UserEmail: userModel.Email, VerifyCode: userModel.VerificationCode}
	return i.messageService.PublishUserCreatedEvent(event)
}

func (i *identityService) LoginUser(loginUserRequest dto.LoginUserRequest) (*dto.TokenResponse, error) {
	userModel, err := i.idendityRepository.GetUserByEmail(loginUserRequest.Email)
	if err != nil {
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(userModel.PasswordHash), []byte(loginUserRequest.Password)); err != nil {
		return nil, err
	}

	return i.createToken(*userModel)
}

func (i *identityService) VerifyUserByCode(code string) error {
	return i.idendityRepository.VerifyUserEmailByCode(code)
}

func (i *identityService) createToken(user model.User) (*dto.TokenResponse, error) {
	expirationTime := time.Now().Add(time.Minute * time.Duration(i.tokenExpirationTime))
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
	tokenString, err := newToken.SignedString([]byte(i.tokenSecretKey))
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

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func generatRandom6digits() string {
	b := make([]byte, 6)
	n, err := io.ReadAtLeast(rand.Reader, b, 6)
	if n != 6 {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

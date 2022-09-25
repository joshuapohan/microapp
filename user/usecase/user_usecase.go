package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	ksuid "github.com/segmentio/ksuid"
	bcrypt "golang.org/x/crypto/bcrypt"

	model "github.com/joshuapohan/microapp/model"
)

type UserUsecase struct {
	userRepository model.UserRepository
}

func NewUserUsecase(userRepository model.UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepository: userRepository,
	}
}

func (u *UserUsecase) GetByKSUID(ctx context.Context, ksuid string) (*model.User, error) {
	return u.userRepository.GetByKSUID(ctx, ksuid)
}

func (u *UserUsecase) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	return u.userRepository.GetByEmail(ctx, email)
}

func (u *UserUsecase) GenerateUserToken(user *model.User) (string, error) {
	var err error
	secret := os.Getenv("JWT_SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ksuid": user.KSUID,
		"email": user.Email,
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Fatal(err)
	}

	return tokenString, nil
}

func (u *UserUsecase) Register(ctx context.Context, email, username, password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", errors.New("Failed to register user")
	}

	newUser := model.User{}
	newUser.KSUID = ksuid.New().String()
	newUser.Username = username
	newUser.Password = string(hash)
	newUser.Email = email
	newUser.CreatedAt = time.Now()

	u.userRepository.Create(newUser)

	token, err := u.GenerateUserToken(&newUser)

	return token, nil
}

func (u *UserUsecase) Login(ctx context.Context, email, password string) (string, error) {
	user, err := u.userRepository.GetByEmail(ctx, email)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	hashedPassword := user.Password
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return "", errors.New("Invalid email/password")
	}
	token, err := u.GenerateUserToken(user)
	if err != nil {
		return "", errors.New("Internal Server Error")
	}
	u.userRepository.CreateLoginHistory(&model.LoginHistory{
		KSUID:     ksuid.New().String(),
		UserId:    user.KSUID,
		CreatedAt: time.Now(),
	})
	return token, nil
}

func (u *UserUsecase) FetchLoginHistories(ctx context.Context, page, perPage int) ([]model.LoginHistory, int64, error) {
	return u.userRepository.FetchLoginHistories(ctx, page, perPage)
}

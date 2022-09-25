package repository

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	model "github.com/joshuapohan/microapp/model"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (ur *UserRepository) Create(u model.User) error {
	res := ur.DB.Create(&u)
	return res.Error
}

func (ur *UserRepository) GetByKSUID(ctx context.Context, ksuid string) (*model.User, error) {
	user := model.User{}
	res := ur.DB.Where("ksuid = ?", ksuid).First(&user)
	if res.Error != nil {
		fmt.Println(user.Email + " ksuid not found")
		return nil, errors.New(user.Email + " ksuid not found")
	}
	return &user, nil
}

func (ur *UserRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	user := model.User{}
	res := ur.DB.Where("email = ?", email).First(&user)
	if res.Error != nil {
		fmt.Println(user.Email + " email not found")
		return nil, errors.New(user.Email + " email not found")
	}
	return &user, nil
}

func (ur *UserRepository) CreateLoginHistory(loginHistory *model.LoginHistory) error {
	res := ur.DB.Create(&loginHistory)
	return res.Error
}

func (ur *UserRepository) FetchLoginHistories(ctx context.Context, page, perPage int) ([]model.LoginHistory, int64, error) {
	histories := make([]model.LoginHistory, 0)
	total := int64(0)

	if user, ok := model.UserFromContext(ctx); ok {
		ur.DB.Where(&model.LoginHistory{UserId: user.KSUID}).Order("created_at desc").Offset((page - 1) * perPage).Limit(perPage).Find(&histories)
		ur.DB.Model(&model.LoginHistory{}).Where("user_id = ?", user.KSUID).Count(&total)
	}
	return histories, total, nil
}

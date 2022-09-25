package model

import (
	"context"
	"time"
)

type User struct {
	KSUID     string    `gorm:"primary_key;column:ksuid" json:"ksuid"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}

type UserUsecase interface {
	GetByKSUID(ctx context.Context, ksuid string) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	Register(ctx context.Context, email, username, password string) (string, error)
	Login(ctx context.Context, email, password string) (string, error)
	FetchLoginHistories(ctx context.Context, page, perPage int) ([]LoginHistory, int64, error)
}

type UserRepository interface {
	Create(user User) error
	GetByKSUID(ctx context.Context, ksuid string) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	CreateLoginHistory(loginHistory *LoginHistory) error
	FetchLoginHistories(ctx context.Context, page, perPage int) ([]LoginHistory, int64, error)
}

// Context related functionalities for storing / retrieving user
type contextKey string

var userContextKey contextKey = "user"

func NewUserContext(ctx context.Context, user *User) context.Context {
	return context.WithValue(ctx, userContextKey, user)
}

func UserFromContext(ctx context.Context) (*User, bool) {
	u, ok := ctx.Value(userContextKey).(*User)
	return u, ok
}

func UserMustFromContext(ctx context.Context) *User {
	u, ok := ctx.Value(userContextKey).(*User)
	if !ok {
		panic("user not found in context")
	}
	return u
}

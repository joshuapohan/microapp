package model

import "time"

type LoginHistory struct {
	KSUID     string    `gorm:"primary_key;column:ksuid" json:"ksuid"`
	UserId    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (LoginHistory) TableName() string {
	return "login_histories"
}

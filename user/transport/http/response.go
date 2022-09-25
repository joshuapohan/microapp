package http

import "github.com/joshuapohan/microapp/model"

type AuthTokenResponse struct {
	Token string `json:"token"`
}

type PaginatedLoginHistoriesResponse struct {
	LoginHistories []model.LoginHistory `json:"login_histories"`
	TotalItems     int64                `json:"total_item"`
}

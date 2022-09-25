package http

import "github.com/joshuapohan/microapp/model"

type FetchAllProvincesResponse struct {
	Provinces []model.Province `json:"provinces"`
}

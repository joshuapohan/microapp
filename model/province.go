package model

import "context"

type SubDistrict struct {
	Name string `json:"name"`
}

type District struct {
	Name         string        `json:"name"`
	SubDistricts []SubDistrict `json:"subdistricts"`
}

type City struct {
	Name      string     `json:"name"`
	Districts []District `json:"district"`
}

type Province struct {
	Name   string `json:"name"`
	Cities []City `json:"cities"`
}

type ProvinceRepository interface {
	FetchAllProvinces(ctx context.Context) ([]Province, error)
}

type ProvinceUsecase interface {
	FetchAllProvinces(ctx context.Context) ([]Province, error)
}

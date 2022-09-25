package usecase

import (
	"context"

	model "github.com/joshuapohan/microapp/model"
)

type ProvinceUsecase struct {
	provinceRepository model.ProvinceRepository
}

func NewProvinceUsecase(provinceRepository model.ProvinceRepository) *ProvinceUsecase {
	return &ProvinceUsecase{provinceRepository: provinceRepository}
}

func (p *ProvinceUsecase) FetchAllProvinces(ctx context.Context) ([]model.Province, error) {
	return p.provinceRepository.FetchAllProvinces(ctx)
}

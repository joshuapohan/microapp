package repository

import (
	"context"

	"github.com/joshuapohan/microapp/model"
)

type MockProvinceRepository struct {
	provincesCache map[string][]model.Province
}

func NewMockProvinceRepository() *MockProvinceRepository {
	return &MockProvinceRepository{
		provincesCache: map[string][]model.Province{},
	}
}

func (p *MockProvinceRepository) FetchAllProvinces(ctx context.Context) ([]model.Province, error) {
	// TO DO : implement cache key based on query parameters
	cacheKey := "all_provinces_with_no_filters" // for mock data use same cache key for all requests
	provinces := make([]model.Province, 0)
	if val, ok := p.provincesCache[cacheKey]; ok {
		return val, nil
	} else {
		provinces = append(provinces, []model.Province{{
			Name: "Province A",
			Cities: []model.City{{
				Name: "City A",
				Districts: []model.District{{
					Name: "District A",
					SubDistricts: []model.SubDistrict{{
						Name: "Subdistrict A",
					}},
				}},
			}},
		},
			{
				Name: "Province B",
				Cities: []model.City{{
					Name: "City B",
					Districts: []model.District{{
						Name: "District B",
						SubDistricts: []model.SubDistrict{{
							Name: "Subdistrict B",
						}},
					}},
				}},
			}}...)
		p.provincesCache[cacheKey] = provinces
	}
	return provinces, nil
}

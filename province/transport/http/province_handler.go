package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	model "github.com/joshuapohan/microapp/model"
	http_util "github.com/joshuapohan/microapp/util"
)

type ProvinceHandler struct {
	provinceUsecase model.ProvinceUsecase
}

func NewProvinceHandler(provinceUsecase model.ProvinceUsecase) *ProvinceHandler {
	return &ProvinceHandler{
		provinceUsecase: provinceUsecase,
	}
}

func (p ProvinceHandler) FetchAllProvinces(w http.ResponseWriter, r *http.Request) {
	provinces, err := p.provinceUsecase.FetchAllProvinces(r.Context())
	if err != nil {
		fmt.Println(err)
		http_util.RespondWithError(w, http.StatusInternalServerError, http_util.Error{Message: "Failed to fetch login histories"})
		return
	}
	res := FetchAllProvincesResponse{
		Provinces: provinces,
	}
	http_util.ResponseJSON(w, res, http.StatusOK)
}

func (p ProvinceHandler) ServeProtected(router *mux.Router) {
	router.HandleFunc("/provinces", p.FetchAllProvinces).Methods("GET")
}

package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	model "github.com/joshuapohan/microapp/model"
	http_util "github.com/joshuapohan/microapp/util"
)

type UserHandler struct {
	userUsecase model.UserUsecase
}

func NewUserHandler(userUsecase model.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

func (u UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	userForm := RegistrationForm{}
	err := json.NewDecoder(r.Body).Decode(&userForm)
	if err != nil {
		fmt.Println(err)
		http_util.RespondWithError(w, http.StatusInternalServerError, http_util.Error{Message: "Failed to decode json"})
		return
	}
	res, err := u.userUsecase.Register(r.Context(), userForm.Email, userForm.Username, userForm.Password)
	if err != nil {
		fmt.Println(err)
		http_util.RespondWithError(w, http.StatusInternalServerError, http_util.Error{Message: err.Error()})
		return
	}
	http_util.ResponseJSON(w, AuthTokenResponse{Token: res}, http.StatusOK)
}

func (u UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	loginForm := LoginForm{}
	err := json.NewDecoder(r.Body).Decode(&loginForm)
	if err != nil {
		fmt.Println(err)
		http_util.RespondWithError(w, http.StatusInternalServerError, http_util.Error{Message: "Failed to decode json"})
		return
	}
	res, err := u.userUsecase.Login(r.Context(), loginForm.Email, loginForm.Password)
	if err != nil {
		fmt.Println(err)
		http_util.RespondWithError(w, http.StatusInternalServerError, http_util.Error{Message: "Failed to register user"})
		return
	}
	http_util.ResponseJSON(w, AuthTokenResponse{Token: res}, http.StatusOK)
}

func (u UserHandler) FetchLoginHistories(w http.ResponseWriter, r *http.Request) {
	paginatedLoginHistoriesForm := PaginatedLoginHistoriesForm{}
	err := json.NewDecoder(r.Body).Decode(&paginatedLoginHistoriesForm)
	if err != nil {
		fmt.Println(err)
		http_util.RespondWithError(w, http.StatusInternalServerError, http_util.Error{Message: "Failed to decode json"})
		return
	}
	histories, total, err := u.userUsecase.FetchLoginHistories(r.Context(), paginatedLoginHistoriesForm.Page, paginatedLoginHistoriesForm.PerPage)
	if err != nil {
		fmt.Println(err)
		http_util.RespondWithError(w, http.StatusInternalServerError, http_util.Error{Message: "Failed to fetch login histories"})
		return
	}
	res := PaginatedLoginHistoriesResponse{
		LoginHistories: histories,
		TotalItems:     total,
	}
	http_util.ResponseJSON(w, res, http.StatusOK)
}

func (u UserHandler) Serve(router *mux.Router) {
	router.HandleFunc("/register", u.Register).Methods("POST", "OPTIONS")
	router.HandleFunc("/login", u.Login).Methods("POST", "OPTIONS")
}

func (u UserHandler) ServeProtected(router *mux.Router) {
	router.HandleFunc("/login-history", u.FetchLoginHistories).Methods("POST", "OPTIONS")
}

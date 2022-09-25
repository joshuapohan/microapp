package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	middleware "github.com/joshuapohan/microapp/app/http/middleware"
	userrepository "github.com/joshuapohan/microapp/user/repository"
	userhandler "github.com/joshuapohan/microapp/user/transport/http"
	userusecase "github.com/joshuapohan/microapp/user/usecase"

	provincerepository "github.com/joshuapohan/microapp/province/repository"
	provincehandler "github.com/joshuapohan/microapp/province/transport/http"
	provinceusecase "github.com/joshuapohan/microapp/province/usecase"
)

func Serve() {
	db, err := gorm.Open(sqlite.Open("microapp.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()

	// Init user module
	userRepository := userrepository.NewUserRepository(db)
	userUsecase := userusecase.NewUserUsecase(userRepository)
	userHandler := userhandler.NewUserHandler(userUsecase)
	userHandler.Serve(router)

	// Init province module
	provinceRepository := provincerepository.NewMockProvinceRepository()
	provinceUsecase := provinceusecase.NewProvinceUsecase(provinceRepository)
	provinceHandler := provincehandler.NewProvinceHandler(provinceUsecase)

	protected := router.PathPrefix("/").Subrouter()
	protected.Use(middleware.AuthMiddleware)
	userHandler.ServeProtected(protected)
	provinceHandler.ServeProtected(protected)

	fmt.Println("listening on port 3000")
	err = http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatal(err)
	}
}

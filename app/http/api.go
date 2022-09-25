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
)

func Serve() {
	db, err := gorm.Open(sqlite.Open("microapp.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()

	userRepository := userrepository.NewUserRepository(db)
	userUsecase := userusecase.NewUserUsecase(userRepository)
	userHandler := userhandler.NewUserHandler(userUsecase)
	userHandler.Serve(router)

	protected := router.PathPrefix("/").Subrouter()
	protected.Use(middleware.AuthMiddleware)
	userHandler.ServeProtected(protected)

	fmt.Println("listening on port 3000")
	err = http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatal(err)
	}
}

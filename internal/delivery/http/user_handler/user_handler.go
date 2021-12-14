package user_handler

import (
	"github.com/gorilla/mux"
	"golang-zazin/domain/usecase"
	"golang-zazin/internal/repository"
	"golang-zazin/internal/usecase/user"
)

type userHandler struct {
	userUC usecase.UserUseCase
}

func UserHandler(r *mux.Router) {
	userRepo := repository.NewUserRepository(nil)
	userInteracts := user.NewUserInteracts(userRepo)
	handler := &userHandler{
		userUC: userInteracts,
	}

	r.HandleFunc("/user", handler.FindList).Methods("GET")
}

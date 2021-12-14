package user

import (
	"golang-zazin/domain/repository"
	"golang-zazin/domain/usecase"
)

type userInteracts struct {
	userRepo repository.UserRepository
}

func NewUserInteracts(userRepo repository.UserRepository) usecase.UserUseCase {
	return &userInteracts{
		userRepo: userRepo,
	}
}

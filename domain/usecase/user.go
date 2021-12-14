package usecase

import (
	"context"
	"golang-zazin/domain/entity"
)

type UserUseCase interface {
	FindUserList(ctx context.Context) ([]*entity.User, error)
}

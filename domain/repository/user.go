package repository

import (
	"context"
	"golang-zazin/domain/entity"
)

type UserRepository interface {
	FindList(ctx context.Context) ([]*entity.User, error)
}

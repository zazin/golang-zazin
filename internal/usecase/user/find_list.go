package user

import (
	"context"
	"golang-zazin/domain/entity"
)

func (u *userInteracts) FindUserList(ctx context.Context) ([]*entity.User, error) {

	users, err := u.userRepo.FindList(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

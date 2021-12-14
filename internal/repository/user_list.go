package repository

import (
	"context"
	"github.com/bxcodec/faker/v3"
	"golang-zazin/domain/entity"
	"time"
)

func (u *userRepository) FindList(ctx context.Context) ([]*entity.User, error) {
	users := make([]*entity.User, 0)

	sum := 0
	for i := 1; i < 5; i++ {
		users = append(users, &entity.User{
			Id:        faker.UUIDDigit(),
			Name:      faker.Name(),
			CreatedAt: time.Now(),
		})
		sum += i
	}
	return users, nil
}

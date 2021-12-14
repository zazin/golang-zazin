package response

import (
	"golang-zazin/domain/entity"
	"time"
)

type UserDetailResponse struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

func MapUserDetailResponse(data *entity.User) UserDetailResponse {
	return UserDetailResponse{
		Id:        data.Id,
		Name:      data.Name,
		CreatedAt: data.CreatedAt,
	}
}

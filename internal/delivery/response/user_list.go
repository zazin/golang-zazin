package response

import (
	"golang-zazin/domain/entity"
)

type UserListResponse struct {
	Users []UserDetailResponse `json:"users"`
}

func MapUserListDomainToResponse(projects []*entity.User) UserListResponse {
	results := make([]UserDetailResponse, 0)
	for _, project := range projects {
		results = append(results, MapUserDetailResponse(project))
	}
	return UserListResponse{
		Users: results,
	}
}

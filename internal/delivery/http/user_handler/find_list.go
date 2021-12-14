package user_handler

import (
	"context"
	"golang-zazin/internal/delivery/response"
	"golang-zazin/pkg/utils"
	"net/http"
)

func (h *userHandler) FindList(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	list, _ := h.userUC.FindUserList(ctx)
	utils.RespondWithJSON(w, http.StatusOK, response.MapUserListDomainToResponse(list))
}

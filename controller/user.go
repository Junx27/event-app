package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"guthub.com/Junx27/event-app/entity"
)

type UserHandler struct {
	repository entity.UserReopository
}

func NewUserHandler(repository entity.UserReopository) *UserHandler {
	return &UserHandler{repository: repository}
}

func (h *UserHandler) GetAllUser(ctx *gin.Context) {
	users, err := h.repository.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": users})
}

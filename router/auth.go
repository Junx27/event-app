package router

import (
	"github.com/gin-gonic/gin"
	"guthub.com/Junx27/event-app/controller"
	"guthub.com/Junx27/event-app/service"
)

func SetupAuthRouter(r *gin.Engine, authService *service.AuthService) {

	authHandler := controller.NewAuthHandler(authService)

	r.POST("/login", authHandler.Login)
	r.POST("/register", authHandler.Register)
	r.POST("/logout", authHandler.Logout)
}

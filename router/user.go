package router

import (
	"github.com/Junx27/event-app/controller"
	"github.com/Junx27/event-app/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserRouter(r *gin.Engine, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	userHandler := controller.NewUserHandler(userRepository)

	userGroup := r.Group("/users")
	{
		userGroup.GET("/", userHandler.GetMany)
	}
}

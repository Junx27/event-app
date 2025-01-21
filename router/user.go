package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"guthub.com/Junx27/event-app/controller"
	"guthub.com/Junx27/event-app/repository"
)

func SetupUserRouter(r *gin.Engine, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	userHandler := controller.NewUserHandler(userRepository)

	userGroup := r.Group("/users")
	{
		userGroup.GET("/", userHandler.GetAllUser)
	}
}

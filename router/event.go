package router

import (
	"github.com/Junx27/event-app/controller"
	"github.com/Junx27/event-app/middleware"
	"github.com/Junx27/event-app/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupEventRouter(r *gin.Engine, db *gorm.DB) {
	eventRepository := repository.NewEventRepository(db)
	eventHandler := controller.NewEventHandler(eventRepository)

	eventGroup := r.Group("/events")
	eventGroup.Use(middleware.AuthProtected(db))
	{
		eventGroup.GET("", eventHandler.GetMany)
		eventGroup.GET("/:id", eventHandler.GetOne)
		eventGroup.POST("", eventHandler.CreateOne)
		eventGroup.PUT("/:id", eventHandler.UpdateOne)
		eventGroup.DELETE("/:id", eventHandler.DeleteOne)
	}
}

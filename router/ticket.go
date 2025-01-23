package router

import (
	"github.com/Junx27/event-app/controller"
	"github.com/Junx27/event-app/middleware"
	"github.com/Junx27/event-app/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupTicketRouter(r *gin.Engine, db *gorm.DB) {
	ticketRepository := repository.NewTicketRepository(db)
	ticketHandler := controller.NewTicketHandler(ticketRepository)

	eventGroup := r.Group("/tickets")
	eventGroup.Use(middleware.AuthProtected(db))
	{
		eventGroup.GET("", ticketHandler.GetMany)
		eventGroup.GET("/:id", ticketHandler.GetOne)
		eventGroup.POST("", ticketHandler.CreateOne)
		eventGroup.PUT("/payment/:id", ticketHandler.UpdateOne)
		eventGroup.DELETE("/:id", ticketHandler.DeleteOne)
	}
}

package router

import (
	"github.com/Junx27/event-app/controller"
	"github.com/Junx27/event-app/middleware"
	"github.com/Junx27/event-app/repository"
	"github.com/Junx27/event-app/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupTicketRouter(r *gin.Engine, db *gorm.DB, ticketService *service.TicketService) {
	ticketRepository := repository.NewTicketRepository(db)
	ticketHandler := controller.NewTicketHandler(ticketRepository, ticketService)

	eventGroup := r.Group("/tickets")
	eventGroup.Use(middleware.AuthProtected(db))
	{
		eventGroup.GET("", ticketHandler.GetMany)
		eventGroup.GET("/:id", ticketHandler.GetOne)
		eventGroup.POST("", ticketHandler.CreateOne)
		eventGroup.PATCH("/payment/:id", ticketHandler.PaymentOne)
		eventGroup.PATCH("/cancel/:id", ticketHandler.CancelOne)
		eventGroup.PATCH("/usage/:id", ticketHandler.UsageTicket)
		eventGroup.DELETE("/:id", ticketHandler.DeleteOne)
	}
}

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
	ticketMiddleware := ticketRepository.(*repository.TicketRepository)

	eventGroup := r.Group("/tickets")
	eventGroup.Use(middleware.AuthProtected(db))
	{
		eventGroup.GET("/booked", middleware.RoleRequired("admin"), ticketHandler.GetManyAdmin)
		eventGroup.GET("", middleware.AccessPermission(ticketMiddleware), ticketHandler.GetMany)
		eventGroup.GET("/:id", middleware.AccessPermission(ticketMiddleware), ticketHandler.GetOne)
		eventGroup.POST("", middleware.RoleRequired("user"), ticketHandler.CreateOne)
		eventGroup.PATCH("/payment/:id", middleware.AccessPermission(ticketMiddleware), middleware.RoleRequired("user"), ticketHandler.PaymentOne)
		eventGroup.PATCH("/cancel/:id", middleware.AccessPermission(ticketMiddleware), middleware.RoleRequired("user"), ticketHandler.CancelOne)
		eventGroup.PATCH("/usage/:id", middleware.AccessPermission(ticketMiddleware), middleware.RoleRequired("user"), ticketHandler.UsageTicket)
		eventGroup.DELETE("/:id", middleware.RoleRequired("admin"), ticketHandler.DeleteOne)
	}
}

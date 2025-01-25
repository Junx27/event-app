package router

import (
	"github.com/Junx27/event-app/controller"
	"github.com/Junx27/event-app/middleware"
	"github.com/Junx27/event-app/service"
	"github.com/gin-gonic/gin"
)

func SetupReportRouter(r *gin.Engine, ticketService *service.TicketService) {
	reportHandler := controller.NewReportHandler(ticketService)

	reportGroup := r.Group("/reports")
	reportGroup.Use(middleware.RoleRequired("admin"))
	{
		reportGroup.GET("/summary", reportHandler.GetSummaryReport)
		reportGroup.GET("/event/:id", reportHandler.GetEventReport)
	}

}

package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Junx27/event-app/helper"
	"github.com/Junx27/event-app/service"
	"github.com/gin-gonic/gin"
)

type ReportHandler struct {
	service *service.TicketService
}

func NewReportHandler(service *service.TicketService) *ReportHandler {
	return &ReportHandler{service: service}
}

func (h *ReportHandler) GetSummaryReport(c *gin.Context) {
	report, err := h.service.GetSummaryReport(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, helper.SuccessResponse("Fetch data successfully", report))
}

func (h *ReportHandler) GetEventReport(c *gin.Context) {
	eventID := c.Param("id")
	id, err := strconv.ParseUint(eventID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	report, err := h.service.GetEventReport(context.Background(), uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, helper.SuccessResponse("Fetch data successfully", report))
}

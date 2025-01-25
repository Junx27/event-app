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

func (h *ReportHandler) GetSummaryReport(ctx *gin.Context) {
	report, err := h.service.GetSummaryReport(context.Background())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, helper.SuccessResponse("Fetch data successfully", report))
}

func (h *ReportHandler) GetEventReport(ctx *gin.Context) {
	eventID := ctx.Param("id")
	id, err := strconv.ParseUint(eventID, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	report, err := h.service.GetEventReport(context.Background(), uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, helper.SuccessResponse("Fetch data successfully", report))
}

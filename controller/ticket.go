package controller

import (
	"math"
	"net/http"
	"strconv"

	"github.com/Junx27/event-app/entity"
	"github.com/Junx27/event-app/helper"
	"github.com/gin-gonic/gin"
)

type TicketHandler struct {
	repository entity.TicketRepository
	service    entity.TicketService
}

func NewTicketHandler(repository entity.TicketRepository, service entity.TicketService) *TicketHandler {
	return &TicketHandler{repository: repository, service: service}
}

func (h *TicketHandler) GetMany(ctx *gin.Context) {
	userID, err := helper.GetUserIDFromCookie(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}
	page := ctx.DefaultQuery("page", "1")
	limit := ctx.DefaultQuery("limit", "10")

	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)

	tickets, totalItems, err := h.repository.GetMany(ctx.Request.Context(), userID, pageInt, limitInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to fetch data"))
		return
	}
	totalPages := int(math.Ceil(float64(totalItems) / float64(limitInt)))

	if pageInt > totalPages {
		pageInt = totalPages
		tickets, _, err = h.repository.GetMany(ctx.Request.Context(), userID, pageInt, limitInt)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to fetch data"))
			return
		}
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Page not found"))
		return
	}

	response := helper.PaginationResponse(tickets, pageInt, limitInt, totalPages, totalItems)
	ctx.JSON(http.StatusOK, helper.SuccessResponse(("Fetch data successfully"), response))
}

func (h *TicketHandler) GetManyAdmin(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")
	limit := ctx.DefaultQuery("limit", "10")

	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)

	tickets, totalItems, err := h.repository.GetManyAdmin(ctx.Request.Context(), pageInt, limitInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to fetch data"))
		return
	}
	totalPages := int(math.Ceil(float64(totalItems) / float64(limitInt)))

	if pageInt > totalPages {
		pageInt = totalPages
		tickets, _, err = h.repository.GetManyAdmin(ctx.Request.Context(), pageInt, limitInt)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to fetch data"))
			return
		}
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Page not found"))
		return
	}

	response := helper.PaginationResponse(tickets, pageInt, limitInt, totalPages, totalItems)
	ctx.JSON(http.StatusOK, helper.SuccessResponse(("Fetch data successfully"), response))
}

func (h *TicketHandler) GetOne(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	ticket, err := h.repository.GetOne(ctx, uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to fetch data"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse(("Fetch data successfully"), ticket))
}

func (h *TicketHandler) CreateOne(ctx *gin.Context) {
	ticket := &entity.Ticket{}
	if err := ctx.ShouldBindJSON(&ticket); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid input"))
		return
	}

	userID, err := helper.GetUserIDFromCookie(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}
	ticket.UserID = userID

	err = h.service.CheckEvent(ctx, ticket.EventID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Quantity not enough"))
		return
	}
	err = h.service.UpdateEvent(ctx, ticket.EventID, ticket.Quantity)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to update event"))
		return
	}
	createTicket, err := h.repository.CreateOne(ctx, ticket)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to create data"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse("Create data successfully", createTicket))
}

func (h *TicketHandler) PaymentOne(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid ticket ID"))
		return
	}
	updateTicket, _, err := h.service.TicketPayment(ctx, uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to fetch data"))
		return
	}
	ctx.JSON(http.StatusOK, helper.SuccessResponse("Update data successfully", updateTicket))
}

func (h *TicketHandler) CancelOne(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid ticket ID"))
		return
	}
	updateTicket, _, err := h.service.TicketCancel(ctx, uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to fetch data"))
		return
	}
	ctx.JSON(http.StatusOK, helper.SuccessResponse("Update data successfully", updateTicket))
}

func (h *TicketHandler) UsageTicket(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid ticket ID"))
		return
	}
	updateTicket, _, err := h.service.TicketUsage(ctx, uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to fetch data"))
		return
	}
	ctx.JSON(http.StatusOK, helper.SuccessResponse("Update data successfully", updateTicket))
}

func (h *TicketHandler) DeleteOne(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	err = h.repository.DeleteOne(ctx, uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to delete data"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse("Delete data successfully", nil))
}

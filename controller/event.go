package controller

import (
	"math"
	"net/http"
	"strconv"

	"github.com/Junx27/event-app/entity"
	"github.com/Junx27/event-app/helper"
	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	repository entity.EventRepository
}

func NewEventHandler(repository entity.EventRepository) *EventHandler {
	return &EventHandler{repository: repository}
}

func (h *EventHandler) GetMany(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")
	limit := ctx.DefaultQuery("limit", "10")

	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)

	events, totalItems, err := h.repository.GetMany(ctx, pageInt, limitInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to fetch data"))
		return
	}
	totalPages := int(math.Ceil(float64(totalItems) / float64(limitInt)))

	if pageInt > totalPages {
		pageInt = totalPages
		events, _, err = h.repository.GetMany(ctx, pageInt, limitInt)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to fetch data"))
			return
		}
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Page not found"))
		return
	}

	pagination := map[string]interface{}{
		"page":        pageInt,
		"limit":       limitInt,
		"total_items": totalItems,
		"total_pages": totalPages,
	}
	responseData := map[string]interface{}{
		"events":     events,
		"pagination": pagination,
	}
	ctx.JSON(http.StatusOK, helper.SuccessResponse(("Fetch data successfully"), responseData))
}

func (h *EventHandler) GetOne(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	event, err := h.repository.GetOne(ctx, uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to fetch data"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse(("Fetch data successfully"), event))
}

func (h *EventHandler) CreateOne(ctx *gin.Context) {
	event := &entity.Event{}
	if err := ctx.ShouldBindJSON(&event); err != nil {
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
	event.UserID = userID
	createdEvent, err := h.repository.CreateOne(ctx, event)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to create data"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse("Create data successfully", createdEvent))
}

func (h *EventHandler) UpdateOne(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	event, err := h.repository.GetOne(ctx, uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to fetch data"))
		return
	}

	updateData := entity.Event{}
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid input"))
		return
	}
	updateFields := map[string]interface{}{
		"id":          event.ID,
		"user_id":     event.UserID,
		"title":       updateData.Title,
		"description": updateData.Description,
		"location":    updateData.Location,
		"date":        updateData.Date,
		"time":        updateData.Time,
		"price":       updateData.Price,
		"quota":       updateData.Quota,
	}

	updatedEvent, err := h.repository.UpdateOne(ctx, uint(id), updateFields)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to update data"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse("Update data successfully", updatedEvent))
}

func (h *EventHandler) DeleteOne(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	err = h.repository.DeleteOne(ctx, uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to delete data"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse("Delete data successfully", nil))
}

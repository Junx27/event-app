package service

import (
	"context"
	"fmt"

	"github.com/Junx27/event-app/entity"
)

type TicketService struct {
	ticketRepo entity.TicketRepository
	eventRepo  entity.EventRepository
}

func NewTicketService(ticketRepo entity.TicketRepository, eventRepo entity.EventRepository) *TicketService {
	return &TicketService{ticketRepo: ticketRepo, eventRepo: eventRepo}
}
func (s *TicketService) CheckEvent(ctx context.Context, id uint) error {
	event, err := s.eventRepo.GetOne(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to retrieve event: %w", err)
	}

	if event.Quota <= 0 {
		return fmt.Errorf("event with ID %d is sold out", id)
	}

	return nil
}

func (s *TicketService) UpdateEvent(ctx context.Context, id uint, qty int) error {
	event, err := s.eventRepo.GetOne(ctx, id)
	if err != nil {
		return err
	}
	quotaAfterUpdate := event.Quota - qty
	status := "available"
	if quotaAfterUpdate == 0 {
		status = "sold out"
	}
	updates := map[string]interface{}{
		"id":          event.ID,
		"title":       event.Title,
		"description": event.Description,
		"location":    event.Location,
		"date":        event.Date,
		"time":        event.Time,
		"price":       event.Price,
		"quota":       event.Quota - qty,
		"status":      status,
		"user_id":     event.UserID,
	}
	_, err = s.eventRepo.UpdateOne(ctx, id, updates)
	if err != nil {
		return err
	}
	return nil
}
func (s *TicketService) TicketCancel(ctx context.Context, id uint) (*entity.Ticket, *entity.TicketResponse, error) {
	ticket, err := s.ticketRepo.GetOne(ctx, id)
	if err != nil {
		return nil, nil, err
	}
	if ticket.Payment {
		return nil, nil, err
	}
	updates := map[string]interface{}{
		"id":       ticket.ID,
		"user_id":  ticket.UserID,
		"event_id": ticket.EventID,
		"quantity": ticket.Quantity,
		"status":   "cancel",
		"payment":  false,
		"usage":    false,
	}
	ticketUpdate, err := s.ticketRepo.UpdateOne(ctx, id, updates)
	if err != nil {
		return nil, nil, err
	}
	ticketResponse := entity.MapToTicketResponse(ticketUpdate)
	return ticketResponse, nil, nil
}
func (s *TicketService) TicketPayment(ctx context.Context, id uint) (*entity.Ticket, *entity.TicketResponse, error) {
	ticket, err := s.ticketRepo.GetOne(ctx, id)
	if err != nil {
		return nil, nil, err
	}
	updates := map[string]interface{}{
		"id":       ticket.ID,
		"user_id":  ticket.UserID,
		"event_id": ticket.EventID,
		"quantity": ticket.Quantity,
		"status":   "paid",
		"payment":  true,
		"usage":    false,
	}
	ticketUpdate, err := s.ticketRepo.UpdateOne(ctx, id, updates)
	if err != nil {
		return nil, nil, err
	}
	ticketResponse := entity.MapToTicketResponse(ticketUpdate)
	return ticketResponse, nil, nil

}
func (s *TicketService) TicketUsage(ctx context.Context, id uint) (*entity.Ticket, *entity.TicketResponse, error) {
	ticket, err := s.ticketRepo.GetOne(ctx, id)
	if err != nil {
		return nil, nil, err
	}
	if !ticket.Payment {
		return nil, nil, err
	}
	updates := map[string]interface{}{
		"id":       ticket.ID,
		"user_id":  ticket.UserID,
		"event_id": ticket.EventID,
		"quantity": ticket.Quantity,
		"status":   ticket.Status,
		"payment":  ticket.Payment,
		"usage":    true,
	}
	ticketUpdate, err := s.ticketRepo.UpdateOne(ctx, id, updates)
	if err != nil {
		return nil, nil, err
	}
	ticketResponse := entity.MapToTicketResponse(ticketUpdate)
	return ticketResponse, nil, nil
}
func (s *TicketService) GetSummaryReport(ctx context.Context) (*entity.SummaryReport, error) {
	tickets, _, err := s.ticketRepo.GetManyAdmin(ctx, 1, 100)
	if err != nil {
		return nil, err
	}

	_, total, err := s.eventRepo.GetMany(ctx, 1, 100)
	if err != nil {
		return nil, err
	}

	var totalSold, totalRevenue int
	for _, ticket := range tickets {
		if ticket.Payment {
			event, err := s.eventRepo.GetOne(ctx, ticket.EventID)
			if err != nil {
				return nil, err
			}

			if event == nil {
				return nil, fmt.Errorf("event with ID %d not found", ticket.EventID)
			}

			totalSold += ticket.Quantity
			totalRevenue += ticket.Quantity * event.Price
		}
	}

	return &entity.SummaryReport{
		TotalTicketsSold: totalSold,
		TotalRevenue:     totalRevenue,
		TotalEvents:      int(total),
	}, nil
}

func (s *TicketService) GetEventReport(ctx context.Context, eventID uint) (*entity.EventReport, error) {
	tickets, _, err := s.ticketRepo.GetManyByEvent(ctx, eventID, 0, 100)
	if err != nil {
		return nil, err
	}

	var totalSold, totalRevenue int
	var eventTitle string
	for _, ticket := range tickets {
		if ticket.EventID == eventID && ticket.Payment {
			event, err := s.eventRepo.GetOne(ctx, ticket.EventID)
			if err != nil {
				return nil, err
			}

			if event == nil {
				return nil, fmt.Errorf("event with ID %d not found", ticket.EventID)
			}

			eventTitle = event.Title
			totalSold += ticket.Quantity
			totalRevenue += ticket.Quantity * event.Price
		}
	}

	return &entity.EventReport{
		EventID:      eventID,
		EventTitle:   eventTitle,
		TotalTickets: totalSold,
		TotalRevenue: totalRevenue,
	}, nil
}

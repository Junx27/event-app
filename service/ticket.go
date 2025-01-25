package service

import (
	"context"

	"github.com/Junx27/event-app/entity"
)

type TicketService struct {
	ticketRepo entity.TicketRepository
	eventRepo  entity.EventRepository
}

func NewTicketService(ticketRepo entity.TicketRepository, eventRepo entity.EventRepository) *TicketService {
	return &TicketService{ticketRepo: ticketRepo, eventRepo: eventRepo}
}
func (s *TicketService) TicketCancel(ctx context.Context, id uint) (*entity.Ticket, *entity.TicketResponse, error) {
	ticket, err := s.ticketRepo.GetOne(ctx, id)
	if err != nil {
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

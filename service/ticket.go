package service

import (
	"context"

	"github.com/Junx27/event-app/entity"
)

type TicketService struct {
	repository entity.TicketRepository
}

func NewTicketService(repository entity.TicketRepository) *TicketService {

	return &TicketService{repository: repository}

}
func (s *TicketService) TicketCancel(ctx context.Context, id uint) (*entity.Ticket, *entity.TicketResponse, error) {
	ticket, err := s.repository.GetOne(ctx, id)
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
	ticketUpdate, err := s.repository.UpdateOne(ctx, id, updates)
	if err != nil {
		return nil, nil, err
	}
	ticketResponse := entity.MapToTicketResponse(ticketUpdate)
	return ticketResponse, nil, nil
}
func (s *TicketService) TicketPayment(ctx context.Context, id uint) (*entity.Ticket, *entity.TicketResponse, error) {
	ticket, err := s.repository.GetOne(ctx, id)
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
	ticketUpdate, err := s.repository.UpdateOne(ctx, id, updates)
	if err != nil {
		return nil, nil, err
	}
	ticketResponse := entity.MapToTicketResponse(ticketUpdate)
	return ticketResponse, nil, nil

}
func (s *TicketService) TicketUsage(ctx context.Context, id uint) (*entity.Ticket, *entity.TicketResponse, error) {
	ticket, err := s.repository.GetOne(ctx, id)
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
	ticketUpdate, err := s.repository.UpdateOne(ctx, id, updates)
	if err != nil {
		return nil, nil, err
	}
	ticketResponse := entity.MapToTicketResponse(ticketUpdate)
	return ticketResponse, nil, nil
}

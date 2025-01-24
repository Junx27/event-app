package repository

import (
	"context"
	"errors"

	"github.com/Junx27/event-app/entity"
	"gorm.io/gorm"
)

type TicketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) entity.TicketRepository {
	return &TicketRepository{db: db}
}

func (r *TicketRepository) GetUserID(id uint) (uint, error) {
	var ticket entity.TicketResponse
	if err := r.db.First(&ticket, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, errors.New("ticket not found")
		}
		return 0, err
	}
	return ticket.UserID, nil
}

func (r *TicketRepository) GetManyByUser(ctx context.Context, userID uint, page, limit int) ([]interface{}, error) {
	tickets, _, err := r.GetMany(ctx, userID, page, limit)
	if err != nil {
		return nil, err
	}
	result := make([]interface{}, len(tickets))
	for i, ticket := range tickets {
		result[i] = ticket
	}
	return result, nil
}

func (r *TicketRepository) GetMany(ctx context.Context, userID uint, page, limit int) ([]*entity.TicketResponse, int64, error) {
	var tickets []*entity.TicketResponse
	var totalItems int64
	if err := r.db.Model(&entity.Ticket{}).Count(&totalItems).Error; err != nil {
		return nil, 0, err
	}
	offset := (page - 1) * limit
	if err := r.db.Offset(offset).Limit(limit).Preload("User").Where("user_id = ?", userID).Find(&tickets).Error; err != nil {
		return nil, 0, err
	}

	return tickets, totalItems, nil
}
func (r *TicketRepository) GetManyAdmin(ctx context.Context, page, limit int) ([]*entity.TicketResponse, int64, error) {
	var tickets []*entity.TicketResponse
	var totalItems int64
	if err := r.db.Model(&entity.Ticket{}).Count(&totalItems).Error; err != nil {
		return nil, 0, err
	}
	offset := (page - 1) * limit
	if err := r.db.Offset(offset).Limit(limit).Preload("User").Find(&tickets).Error; err != nil {
		return nil, 0, err
	}

	return tickets, totalItems, nil
}

func (r *TicketRepository) GetOne(ctx context.Context, id uint) (*entity.TicketResponse, error) {
	ticket := &entity.TicketResponse{}
	if res := r.db.Model(ticket).Where("id = ?", id).Preload("User").First(ticket); res.Error != nil {
		return nil, res.Error
	}

	return ticket, nil
}

func (r *TicketRepository) CreateOne(ctx context.Context, ticket *entity.Ticket) (*entity.Ticket, error) {
	if err := r.db.WithContext(ctx).Create(ticket).Error; err != nil {
		return nil, err
	}

	return ticket, nil
}

func (r *TicketRepository) UpdateOne(ctx context.Context, id uint, updateData map[string]interface{}) (*entity.Ticket, error) {
	ticket := &entity.Ticket{}
	res := r.db.Model(&ticket).Where("id = ?", id).Updates(updateData)

	if res.Error != nil {
		return nil, res.Error
	}

	return ticket, nil
}

func (r *TicketRepository) DeleteOne(ctx context.Context, id uint) error {
	ticket := &entity.Ticket{}
	res := r.db.Model(&ticket).Where("id = ?", id).Delete(ticket)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

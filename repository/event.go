package repository

import (
	"context"

	"github.com/Junx27/event-app/entity"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) entity.EventRepository {
	return &EventRepository{db: db}
}

func (r *EventRepository) GetMany(ctx context.Context, page, limit int, nameFilter, locationFilter, categoryFilter string) ([]*entity.EventResponse, int64, error) {
	var events []*entity.EventResponse
	var totalItems int64
	query := r.db.Model(&entity.Event{})

	if nameFilter != "" {
		query = query.Where("title LIKE ?", "%"+nameFilter+"%")
	}

	if locationFilter != "" {
		query = query.Where("location LIKE ?", "%"+locationFilter+"%")
	}

	if categoryFilter != "" {
		query = query.Where("category LIKE ?", "%"+categoryFilter+"%")
	}

	if err := query.Count(&totalItems).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := query.Offset(offset).Limit(limit).Preload("User").Find(&events).Error; err != nil {
		return nil, 0, err
	}

	return events, totalItems, nil
}

func (r *EventRepository) GetOne(ctx context.Context, id uint) (*entity.EventResponse, error) {
	event := &entity.EventResponse{}
	if res := r.db.Model(event).Where("id = ?", id).Preload("User").First(event); res.Error != nil {
		return nil, res.Error
	}

	return event, nil
}

func (r *EventRepository) CreateOne(ctx context.Context, event *entity.Event) (*entity.Event, error) {
	if err := r.db.WithContext(ctx).Create(event).Error; err != nil {
		return nil, err
	}

	return event, nil
}

func (r *EventRepository) UpdateOne(ctx context.Context, id uint, updateData map[string]interface{}) (*entity.Event, error) {
	event := &entity.Event{}
	res := r.db.Model(&event).Where("id = ?", id).Updates(updateData)

	if res.Error != nil {
		return nil, res.Error
	}

	return event, nil
}

func (r *EventRepository) DeleteOne(ctx context.Context, id uint) error {
	event := &entity.Event{}
	res := r.db.Model(&event).Where("id = ?", id).Delete(event)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

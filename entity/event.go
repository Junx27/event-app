package entity

import "context"

type BaseModelEvent struct{}

func (BaseModelEvent) TableName() string {
	return "events"
}

type SummaryReport struct {
	TotalTicketsSold int `json:"total_tickets_sold"`
	TotalRevenue     int `json:"total_revenue"`
	TotalEvents      int `json:"total_events"`
}
type EventReport struct {
	EventID      uint   `json:"event_id"`
	EventTitle   string `json:"event_title"`
	TotalTickets int    `json:"total_tickets"`
	TotalRevenue int    `json:"total_revenue"`
}

type Event struct {
	BaseModelEvent
	ID          uint         `json:"id" gorm:"primaryKey"`
	UserID      uint         `json:"user_id" gorm:"not null"`
	Title       string       `json:"title" gorm:"not null"`
	Category    string       `json:"category" gorm:"not null"`
	Description string       `json:"description" gorm:"not null"`
	Location    string       `json:"location" gorm:"not null"`
	Date        string       `json:"date" gorm:"not null"`
	Time        string       `json:"time" gorm:"not null"`
	Price       int          `json:"price" gorm:"not null"`
	Quota       int          `json:"quota" gorm:"not null"`
	Status      string       `json:"status" gorm:"default:not yet"`
	User        UserResponse `json:"-" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

type EventResponse struct {
	BaseModelEvent
	ID          uint         `json:"id" gorm:"primaryKey"`
	UserID      uint         `json:"user_id" gorm:"not null"`
	Title       string       `json:"title" gorm:"not null"`
	Category    string       `json:"category" gorm:"not null"`
	Description string       `json:"description" gorm:"not null"`
	Location    string       `json:"location" gorm:"not null"`
	Date        string       `json:"date" gorm:"not null"`
	Time        string       `json:"time" gorm:"not null"`
	Price       int          `json:"price" gorm:"not null"`
	Quota       int          `json:"quota" gorm:"not null"`
	Status      string       `json:"status" gorm:"default:not yet"`
	User        UserResponse `json:"user" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

type EventRepository interface {
	GetMany(ctx context.Context, page, limit int, nameFilter, locationFilter, categoryFilter string) ([]*EventResponse, int64, error)
	GetOne(ctx context.Context, id uint) (*EventResponse, error)
	CreateOne(ctx context.Context, event *Event) (*Event, error)
	UpdateOne(ctx context.Context, id uint, updateData map[string]interface{}) (*Event, error)
	DeleteOne(ctx context.Context, id uint) error
}

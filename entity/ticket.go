package entity

import "context"

type BaseModelTiket struct{}

func (BaseModelTiket) TableName() string {
	return "tickets"
}

type Ticket struct {
	BaseModelTiket
	ID      uint         `json:"id" gorm:"primaryKey"`
	UserID  uint         `json:"user_id" gorm:"not null"`
	EventID uint         `json:"event_id" gorm:"not null"`
	Payment bool         `json:"payment" gorm:"not null"`
	Usage   bool         `json:"usage" gorm:"not null"`
	User    UserResponse `json:"-" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

type TicketResponse struct {
	BaseModelTiket
	ID       uint         `json:"id" gorm:"primaryKey"`
	UserID   uint         `json:"user_id" gorm:"not null"`
	EventID  uint         `json:"event_id" gorm:"not null"`
	Quantity int          `json:"quantity" gorm:"not null"`
	Payment  bool         `json:"payment" gorm:"default:false"`
	Usage    bool         `json:"usage" gorm:"default:false"`
	User     UserResponse `json:"user" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

type TicketRepository interface {
	GetMany(ctx context.Context, page, limit int) ([]*TicketResponse, int64, error)
	GetOne(ctx context.Context, id uint) (*TicketResponse, error)
	CreateOne(ctx context.Context, ticket *Ticket) (*Ticket, error)
	UpdateOne(ctx context.Context, id uint, updateData map[string]interface{}) (*Ticket, error)
	DeleteOne(ctx context.Context, id uint) error
}

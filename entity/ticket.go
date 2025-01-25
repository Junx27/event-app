package entity

import "context"

type BaseModelTiket struct{}

func (BaseModelTiket) TableName() string {
	return "tickets"
}

type Ticket struct {
	BaseModelTiket
	ID       uint         `json:"id" gorm:"primaryKey"`
	UserID   uint         `json:"user_id" gorm:"not null"`
	EventID  uint         `json:"event_id" gorm:"not null"`
	Quantity int          `json:"quantity" gorm:"not null"`
	Status   string       `json:"status" gorm:"default:pending"`
	Payment  bool         `json:"payment" gorm:"default:false"`
	Usage    bool         `json:"usage" gorm:"default:false"`
	User     UserResponse `json:"-" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Event    Event        `json:"-" gorm:"foreignKey:EventID;constraint:OnDelete:CASCADE"`
}

type TicketResponse struct {
	BaseModelTiket
	ID       uint         `json:"id" gorm:"primaryKey"`
	UserID   uint         `json:"user_id" gorm:"not null"`
	EventID  uint         `json:"event_id" gorm:"not null"`
	Quantity int          `json:"quantity" gorm:"not null"`
	Status   string       `json:"status" gorm:"default:pending"`
	Payment  bool         `json:"payment" gorm:"default:false"`
	Usage    bool         `json:"usage" gorm:"default:false"`
	User     UserResponse `json:"user" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Event    Event        `json:"event" gorm:"foreignKey:EventID;constraint:OnDelete:CASCADE"`
}

type TicketRepository interface {
	GetManyAdmin(ctx context.Context, page, limit int) ([]*TicketResponse, int64, error)
	GetMany(ctx context.Context, userId uint, page, limit int) ([]*TicketResponse, int64, error)
	GetManyByEvent(ctx context.Context, eventId uint, page, limit int) ([]*TicketResponse, int64, error)
	GetOne(ctx context.Context, id uint) (*TicketResponse, error)
	CreateOne(ctx context.Context, ticket *Ticket) (*Ticket, error)
	UpdateOne(ctx context.Context, id uint, updateData map[string]interface{}) (*Ticket, error)
	DeleteOne(ctx context.Context, id uint) error
}

type TicketService interface {
	CheckEvent(ctx context.Context, id uint) error
	UpdateEvent(ctx context.Context, id uint, qty int) error
	TicketCancel(ctx context.Context, id uint) (*Ticket, *TicketResponse, error)
	TicketPayment(ctx context.Context, id uint) (*Ticket, *TicketResponse, error)
	TicketUsage(ctx context.Context, id uint) (*Ticket, *TicketResponse, error)
}

func MapToTicketResponse(ticket *Ticket) *Ticket {

	return &Ticket{
		ID:       ticket.ID,
		UserID:   ticket.UserID,
		EventID:  ticket.EventID,
		Quantity: ticket.Quantity,
		Status:   ticket.Status,
		Payment:  ticket.Payment,
		Usage:    ticket.Usage,
	}

}

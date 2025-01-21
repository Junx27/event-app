package entity

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string `gorm:"not null"`
}

type UserReopository interface {
	GetAll() ([]*User, error)
}

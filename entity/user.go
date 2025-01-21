package entity

type User struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	UserName    string `json:"username" gorm:"unique;not null"`
	Email       string `json:"email" gorm:"unique;not null"`
	Password    string `json:"password" gorm:"not null"`
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role" gorm:"default:user"`
}

type UserReopository interface {
	GetMany() ([]*User, error)
}

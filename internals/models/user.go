package models

type User struct {
	ID        string    `json:"id" gorm:"unique;primaryKey" validate:"required"`
	Name      string    `json:"name" gorm:"not null" validate:"required,min=3,max=100"`
	Email     string    `json:"email" gorm:"not null" validate:"required,email"`
	Phone     string    `json:"phone_number" gorm:"not null" validate:"required"`
	Addresses []Address `json:"addresses" gorm:"foreignKey:UserID"`
}

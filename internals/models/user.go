package models

type User struct {
	ID        string    `json:"id" gorm:"unique;primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"not null"`
	Phone     string    `json:"phone_number" gorm:"not null"`
	Addresses []Address `json:"addresses" gorm:"foreignKey:UserID"`
}

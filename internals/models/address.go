package models

type Address struct {
	ID      uint   `gorm:"unique;primaryKey:autoIncrement"`
	City    string `json:"city" gorm:"not null" validate:"required,min=3,max=100"`
	State   string `json:"state" gorm:"not null" validate:"required,min=3,max=100"`
	Zipcode string `json:"zip_code" gorm:"not null" validate:"required,min=3,max=100"`
	Country string `json:"country" gorm:"not null" validate:"required,min=3,max=100"`
	UserID  string
}

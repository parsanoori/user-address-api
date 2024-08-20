package models

type Address struct {
	ID      uint   `gorm:"unique;primaryKey:autoIncrement"`
	City    string `json:"city" gorm:"not null"`
	State   string `json:"state" gorm:"not null"`
	Zipcode string `json:"zip_code" gorm:"not null"`
	Country string `json:"country" gorm:"not null"`
	UserID  string
}

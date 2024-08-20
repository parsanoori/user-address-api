package main

type User struct {
	ID        string    `json:"id" gorm:"unique;primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"not null"`
	Phone     string    `json:"phone_number" gorm:"not null"`
	Addresses []Address `json:"addresses" gorm:"foreignKey:UserID"`
}

type Address struct {
	ID      uint   `gorm:"unique;primaryKey:autoIncrement"`
	City    string `json:"city" gorm:"not null"`
	State   string `json:"state" gorm:"not null"`
	Zipcode string `json:"zip_code" gorm:"not null"`
	Country string `json:"country" gorm:"not null"`
	UserID  string
}

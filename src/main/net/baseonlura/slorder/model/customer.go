package model

import "time"

type Customer struct {
	CustomerId     uint `gorm:"primaryKey"`
	CustomerName   string
	PostNumber     string
	Prefectures    string
	Municipalities string
	Address        string
	BuildingName   string
	PhoneNumber    string
	MailAddress    string
	CreateDate     string `gorm:"column:customer_create_date"`
	Note           string
	LastUpdateDate time.Time
}

func (Customer) TableName() string {
	return "customer"
}

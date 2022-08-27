package model

import "time"

type User struct {
	UserId         uint `gorm:"primaryKey"`
	Account        string
	LastName       string
	FirstName      string
	CreateDate     string
	InvalidDate    string
	BuiltinFlag    bool
	LastUpdateDate time.Time
}

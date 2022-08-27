package model

import "time"

type Member struct {
	MemberId       uint `gorm:"primaryKey"`
	LastName       string
	FirstName      string
	PostNumber     string
	Prefectures    string
	Municipalities string
	Address        string
	BuildingName   string
	PhoneNumber    string
	MailAddress    string
	JoinDate       time.Time
	RetiredDate    time.Time
	RetiredReason  string
	UnitCost       int
	LastUpdateDate time.Time `gorm:"autoUpdateTime"`
}

func (Member) TableName() string {
	return "member"
}

func (m Member) GetFullName() string {
	return m.LastName + " " + m.FirstName
}

package model

import "time"

type OtherCostKind struct {
	KindId   int    `gorm:"primaryKey" json:"kindId"`
	KindName string `json:"name"`
}

func (OtherCostKind) TableName() string {
	return "cost_kind"
}

type OtherCost struct {
	CostId    uint `gorm:"primaryKey"`
	ProjectId uint `gorm:"foreignKey:ProjectId"`
	CostName  string
	KindId    int
	CostKind  OtherCostKind `gorm:"foreignKey:KindId"`
	BuyDate   time.Time
	Cost      int64
}

func (OtherCost) TableName() string {
	return "cost"
}

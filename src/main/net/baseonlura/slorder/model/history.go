package model

import (
	"strings"
	"time"
)

type CommonOperationKind interface {
	IntKey() int
	NewCreateKey() int
	SaveKey() int
}

type ProjectOperationKind int

const (
	RECEIVE ProjectOperationKind = iota + 1
	DELIVERY
	ACCEPTANCE
	PEYMENTED
	LOST_ORDER
	RECEIVE_CANCEL
	DELIVERY_CANCEL
	ACCEPTANCE_CANCEL
	PAYMENTED_CANCEL
	LOST_ORDER_CANCEL
)

var ProjectOperationKindMap = map[ProjectOperationKind]string{
	RECEIVE:           "receive",
	DELIVERY:          "delivery",
	ACCEPTANCE:        "acceptance",
	PEYMENTED:         "peymented",
	LOST_ORDER:        "lost_order",
	RECEIVE_CANCEL:    "receive_cancel",
	DELIVERY_CANCEL:   "delivery_cancel",
	ACCEPTANCE_CANCEL: "acceptance_cancel",
	PAYMENTED_CANCEL:  "paymented_cancel",
	LOST_ORDER_CANCEL: "lost_order_cancel",
}

type ProjectHistory struct {
	ProjectId     uint `gorm:"primaryKey"`
	ProjectName   string
	UpdateUserId  uint `gorm:"primaryKey;column:user_account"`
	UpdateUser    User `gorm:"foreignKey:UpdateUserId"`
	LastName      string
	FirstName     string
	OperationDate time.Time `gorm:"primaryKey"`
	OperationKind int
}

func (p ProjectOperationKind) String() string {
	s, ok := ProjectOperationKindMap[p]
	if ok {
		return s
	}
	return ""
}
func (p ProjectOperationKind) IntKey() int {
	return int(p)
}

func (ProjectOperationKind) SaveKey() int {
	key, _ := commonSave()
	return key
}

func (ProjectOperationKind) NewCreateKey() int {
	key, _ := commonNewCrate()
	return key
}

func (ProjectOperationKind) GetProjectHistoryByKey(k int) ProjectOperationKind {
	for key := range ProjectOperationKindMap {
		if key.IntKey() == k {
			return key
		}
	}
	return -1
}

func (p ProjectOperationKind) GetKinds(s string) ProjectOperationKind {
	for key, value := range ProjectOperationKindMap {
		if strings.EqualFold(s, value) {
			return key
		}
	}
	return -1
}

func (p ProjectHistory) GetFullName() string {
	return p.LastName + " " + p.FirstName
}

func (ProjectHistory) TableName() string {
	return "project_operation_history"
}

func commonSave() (int, string) {
	return 100, "save"
}

func commonNewCrate() (int, string) {
	return 0, "new_create"
}

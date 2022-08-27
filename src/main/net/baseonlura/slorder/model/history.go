package model

import (
	"strings"
)

type ProjectOperationKind int

const (
	NEW_CREATE ProjectOperationKind = iota
	RECEIVE
	DELIVERY
	ACCEPTANCE
	PEYMENTED
	LOST_ORDER
	RECEIVE_CANCEL
	DELIVERY_CANCEL
	ACCEPTANCE_CANCEL
	LOST_ORDER_CANCEL
)

var ProjectOperationKindMap = map[ProjectOperationKind]string{
	NEW_CREATE:        "new_create",
	RECEIVE:           "receive",
	DELIVERY:          "delivery",
	ACCEPTANCE:        "acceptance",
	PEYMENTED:         "peymented",
	LOST_ORDER:        "lost_order",
	RECEIVE_CANCEL:    "receive_cancel",
	DELIVERY_CANCEL:   "delivery_cancel",
	ACCEPTANCE_CANCEL: "acceptance_cancel",
	LOST_ORDER_CANCEL: "lost_order_cancel",
}

type ProjectHistory struct {
	ProjectId     uint `gorm:"primaryKey"`
	ProjectName   string
	UpdateUserId  uint `gorm:"primaryKey"`
	UpdateUser    User `gorm:"foreignKey:UpdateUserId"`
	LastName      string
	FirstName     string
	OperationDate string `gorm:"primaryKey"`
	OperationKind ProjectOperationKind
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

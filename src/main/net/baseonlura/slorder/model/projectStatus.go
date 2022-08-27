package model

import "strings"

type ProjectStatus int

const (
	ESTIMATION ProjectStatus = iota + 1
	RECEIVED
	DELIVERED
	ACCEPTANCED
	PAYMENTED
	LOSTORDER
)

var projectStatusMap = map[ProjectStatus]string{
	ESTIMATION:  "estimation",
	RECEIVED:    "received",
	DELIVERED:   "delivered",
	ACCEPTANCED: "acceptanced",
	PAYMENTED:   "paymented",
	LOSTORDER:   "lostOrder",
}

func (p ProjectStatus) String() string {
	s, ok := projectStatusMap[p]
	if ok {
		return s
	}
	return ""
}

func (ProjectStatus) GetProjectStatus(s string) ProjectStatus {
	for key, value := range projectStatusMap {
		if strings.EqualFold(s, value) {
			return key
		}
	}
	return -1
}

func (ProjectStatus) GetProjectStatusByKey(k int) ProjectStatus {
	for key := range projectStatusMap {
		if key.IntKey() == k {
			return key
		}
	}
	return -1
}

func (p ProjectStatus) IntKey() int {
	return int(p)
}

func (p ProjectStatus) IsUnfinished() bool {
	switch p {
	case ESTIMATION:
		return true
	case RECEIVED:
		return true
	case DELIVERED:
		return true
	case ACCEPTANCED:
		return true
	case PAYMENTED:
		return false
	case LOSTORDER:
		return false
	default:
		return false
	}
}

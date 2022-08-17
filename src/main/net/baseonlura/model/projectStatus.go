package model

import "strings"

type ProjectStatus int

const (
	UNKNOWN ProjectStatus = iota
	ESTIMATION
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
	return "unknown"
}

func getProjectStatus(s string) ProjectStatus {
	for key, value := range projectStatusMap {
		if strings.EqualFold(s, value) {
			return key
		}
	}
	return UNKNOWN
}

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

func GetProjectStatusByKey(k int) ProjectStatus {
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

func (before ProjectStatus) IsEstimationToReceived(after ProjectStatus) bool {
	if before == ESTIMATION && after == RECEIVED {
		return true
	}
	return false
}
func (before ProjectStatus) IsEstimationToLostorder(after ProjectStatus) bool {
	if before == ESTIMATION && after == LOSTORDER {
		return true
	}
	return false
}
func (before ProjectStatus) IsReceivedToDelivered(after ProjectStatus) bool {
	if before == RECEIVED && after == DELIVERED {
		return true
	}
	return false
}
func (before ProjectStatus) IsReceivedToLostorder(after ProjectStatus) bool {
	if before == RECEIVED && after == LOSTORDER {
		return true
	}
	return false
}
func (before ProjectStatus) IsDeliveredToAcceptanced(after ProjectStatus) bool {
	if before == DELIVERED && after == ACCEPTANCED {
		return true
	}
	return false
}
func (before ProjectStatus) IsAcceptancedToPaymented(after ProjectStatus) bool {
	if before == ACCEPTANCED && after == PAYMENTED {
		return true
	}
	return false
}
func (before ProjectStatus) IsPaymentedToAcceptanced(after ProjectStatus) bool {
	if before == PAYMENTED && after == ACCEPTANCED {
		return true
	}
	return false
}
func (before ProjectStatus) IsAcceptancedToDelivered(after ProjectStatus) bool {
	if before == ACCEPTANCED && after == DELIVERED {
		return true
	}
	return false
}
func (before ProjectStatus) IsDeliveredToReceived(after ProjectStatus) bool {
	if before == DELIVERED && after == RECEIVED {
		return true
	}
	return false
}
func (before ProjectStatus) IsReceivedToEstimation(after ProjectStatus) bool {
	if before == RECEIVED && after == ESTIMATION {
		return true
	}
	return false
}
func (before ProjectStatus) IsLostorderToReceived(after ProjectStatus) bool {
	if before == LOSTORDER && after == RECEIVED {
		return true
	}
	return false
}
func (before ProjectStatus) IsLostorderToEstimation(after ProjectStatus) bool {
	if before == LOSTORDER && after == ESTIMATION {
		return true
	}
	return false
}

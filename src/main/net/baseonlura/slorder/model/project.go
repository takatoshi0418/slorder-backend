package model

type Work struct {
	WorkDate string
	WorkTime float32
}

type ProjectMember struct {
	Member     Member
	AssignDate string
	RejectDate string
	UnitCost   int
	Works      []Work
}

type Project struct {
	ProjectNo             string
	ProjectName           string
	Customer              Customer
	Status                ProjectStatus
	StartDate             string
	LimitDate             string
	EstimateOpeWorkByTime float32
	EstimateOperatingCost int64
	EstimateOtherCost     int64
	CreateUser            User
	CreateDate            string
	OrderUser             User
	OrderDate             string
	ReceiveAmount         int64
	DeliveryUser          User
	DeliveryDate          string
	AcceptanceLimitDate   string
	AcceptanceUser        User
	AcceptanceDate        string
	PaymentLimitDate      string
	PaymentUser           User
	PaymentDate           string
	LostOrderDate         string
	ProjectMembers        []ProjectMember
	OtherCosts            []OtherCost
	ProjectHistories      []ProjectHistory
}

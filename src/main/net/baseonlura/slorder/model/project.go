package model

type PaymentData struct {
	OparatingWorkByTime float32
	OparatingCost       int64
	OtherCost           int64
	Proceeds            int64
}

type Payment struct {
	Estimate PaymentData
	Actual   PaymentData
}

type ProjectMenber struct {
	Member     Member
	AssignDate string
	RejectDate string
	UnitCost   int
}

type Project struct {
	ProjectNo            string
	Name                 string
	Customer             Customer
	Status               ProjectStatus
	StartDate            string
	LimitDate            string
	Payment              Payment
	CreateUser           User
	CreateDate           string
	OrderUser            User
	OrderDate            string
	ReceiveAmount        int64
	DeliveryUser         User
	DeliveryDate         string
	AcceptancedLimitDate string
	AcceptancedUser      User
	AcceptancedDate      string
	PaymentLimitDate     string
	PaymentedUser        User
	PaymentedDate        string
	LostOrderDate        string
	ProjectMember        []ProjectMenber
	OtherCost            []OtherCost
	ProjectHistory       []ProjectHistory
}

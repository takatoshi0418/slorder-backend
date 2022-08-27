package model

import "time"

type Project struct {
	SimpleProject    SimpleProject `gorm:"embedded"`
	CreateUser       User          `gorm:"foreignKey:UserId"`
	CreateDate       time.Time     `gorm:"autoCreateTime"`
	OrderUser        User          `gorm:"foreignKey:UserId"`
	OrderDate        time.Time
	DeliveryUser     User `gorm:"foreignKey:UserId"`
	DeliveryDate     time.Time
	AcceptancedUser  User `gorm:"foreignKey:UserId"`
	AcceptancedDate  time.Time
	PaymentedUser    User `gorm:"foreignKey:UserId"`
	PaymentedDate    time.Time
	ProjectMembers   []ProjectMember  `gorm:"foreignKey:ProjectId"`
	OtherCosts       []OtherCost      `gorm:"foreignKey:CostId"`
	Works            []Work           `gorm:"foreignKey:ProjectId"`
	ProjectHistories []ProjectHistory `gorm:"foreignKey:ProjectId"`
}

func (Project) TableName() string {
	return "project"
}

type SimpleProject struct {
	ProjectId             uint `gorm:"primaryKey"`
	ProjectName           string
	Customer              Customer `gorm:"foreignKey:CustomerId"`
	ProjectStatus         int
	StartDate             time.Time
	LimitDate             time.Time
	EstimateOperatingTime *float32
	EstimateOperatingCost *int64
	OthersCostAmount      *int64
	ReceiveAmount         int64 `gorm:"column:order_amount"`
	AcceptanceLimitDate   time.Time
	PaymentLimitDate      time.Time
	LostOrderDate         time.Time
	LastUpdateDate        time.Time `gorm:"autoUpdateTime"`
}

func (s SimpleProject) GetStatus() ProjectStatus {
	return ProjectStatus(s.ProjectStatus)
}

func (SimpleProject) TableName() string {
	return "project"
}

func (s SimpleProject) GetEstimateOperatingTime() float32 {
	if s.EstimateOperatingTime == nil {
		return -1
	}
	return *s.EstimateOperatingTime
}

func (s SimpleProject) GetEstimateOperatingCost() int64 {
	if s.EstimateOperatingCost == nil {
		return -1
	}
	return *s.EstimateOperatingCost
}

func (s SimpleProject) GetOthersCostAmount() int64 {
	if s.OthersCostAmount == nil {
		return -1
	}
	return *s.OthersCostAmount
}

type ProjectMember struct {
	ProjectId  uint      `gorm:"primaryKey"`
	MemberId   uint      `gorm:"primaryKey"`
	Member     Member    `gorm:"foreignKey:MemberId"`
	AssignDate time.Time `gorm:"primaryKey"`
	RejectDate time.Time
	UnitCost   int
	// this column is use only Operating register function
	Work Work `gorm:"-:all"`
}

func (ProjectMember) TableName() string {
	return "project_member"
}

type Work struct {
	ProjectId uint      `gorm:"primaryKey;foreignKey:ProjectId"`
	MemberId  uint      `gorm:"primaryKey"`
	Member    Member    `gorm:"foreignKey:MemberId"`
	WorkDate  time.Time `gorm:"primaryKey"`
	WorkTime  int
}

func (Work) TableName() string {
	return "work"
}

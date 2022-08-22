package viewModel

import (
	"logicApi/src/main/net/baseonlura/slorder/model"
)

type ProjectItem struct {
	Status           int              `json:"status"`
	BasicInfo        BasicInfo        `json:"basic"`
	Payment          Payment          `json:"payment"`
	Members          []ProjectMember  `json:"members"`
	OtherCosts       []OtherCost      `json:"otherCosts"`
	ProjectHistories []ProjectHistory `json:"histories"`
}

type BasicInfo struct {
	ProjectNo     string `json:"no"`
	ProjectName   string `json:"name"`
	Client        string `json:"client"`
	StartDate     string `json:"startDate"`
	LimitDate     string `json:"limitDate"`
	ReceiveAmount int64  `json:"receiveAmount"`
}
type Payment struct {
	OperatingWorkByTime float32 `json:"operatingWorkByTime"`
	OperatingCost       int64   `json:"operatingCost"`
	OtherCost           int64   `json:"otherCost"`
}
type ProjectMember struct {
	MemberID      string  `json:"value"`
	UnitCost      int     `json:"unit"`
	OperatingTime float32 `json:"operatingTime"`
}
type OtherCost struct {
	Name    string `json:"name"`
	Kind    int    `json:"kind"`
	BuyDate string `json:"buyDate"`
	Price   int64  `json:"price"`
}
type ProjectHistory struct {
	Name string `json:"name"`
	Date string `json:"date"`
	Kind int    `json:"kind"`
}

func (pvm *ProjectItem) ToViewModel(p model.Project) {
	pvm.Status = p.Status.IntKey()
	// basicInfo
	basicInfo := new(BasicInfo)
	basicInfo.ProjectNo = p.ProjectNo
	basicInfo.ProjectName = p.ProjectName
	basicInfo.Client = p.Customer.CustomerId
	basicInfo.StartDate = p.StartDate
	basicInfo.LimitDate = p.LimitDate
	basicInfo.ReceiveAmount = p.ReceiveAmount

	// payment
	payment := new(Payment)
	payment.OperatingWorkByTime = p.EstimateOpeWorkByTime
	payment.OperatingCost = p.EstimateOperatingCost
	payment.OtherCost = p.EstimateOtherCost

	// members
	var members []ProjectMember
	for _, pm := range p.ProjectMembers {
		member := new(ProjectMember)
		member.MemberID = pm.Member.MemberId
		member.UnitCost = pm.UnitCost
		member.OperatingTime = 0
		for _, w := range pm.Works {
			member.OperatingTime += w.WorkTime
		}
		members = append(members, *member)
	}

	// otherCosts
	var otherCosts []OtherCost
	for _, oc := range p.OtherCosts {
		otherCost := new(OtherCost)
		otherCost.Name = oc.Name
		otherCost.Kind = oc.CostKind.KindId
		otherCost.BuyDate = oc.BuyDate
		otherCost.Price = oc.Cost
		otherCosts = append(otherCosts, *otherCost)
	}

	// histories
	var histories []ProjectHistory
	for _, ph := range p.ProjectHistories {
		history := new(ProjectHistory)
		history.Name = ph.GetFullName()
		history.Date = ph.OperationDate
		history.Kind = ph.OperationKind.IntKey()
		histories = append(histories, *history)
	}

	pvm.BasicInfo = *basicInfo
	pvm.Payment = *payment
	pvm.Members = members
	pvm.OtherCosts = otherCosts
	pvm.ProjectHistories = histories
}

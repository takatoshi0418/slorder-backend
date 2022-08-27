package viewModel

import (
	"logicApi/src/main/net/baseonlura/slorder/model"
)

type ProjectItem struct {
	Status           int                       `json:"status"`
	BasicInfo        BasicInfo                 `json:"basic"`
	Payment          Payment                   `json:"payment"`
	Members          []ProjectMemberExtOpeTime `json:"members"`
	OtherCosts       []OtherCost               `json:"otherCosts"`
	ProjectHistories []ProjectHistory          `json:"histories"`
}

type BasicInfo struct {
	ProjectNo     uint   `json:"no"`
	ProjectName   string `json:"name"`
	Client        uint   `json:"client"`
	StartDate     string `json:"startDate"`
	LimitDate     string `json:"limitDate"`
	ReceiveAmount int64  `json:"receiveAmount"`
}
type Payment struct {
	OperatingWorkByTime float32 `json:"operatingWorkByTime"`
	OperatingCost       int64   `json:"operatingCost"`
	OtherCost           int64   `json:"otherCost"`
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
	pvm.Status = p.SimpleProject.ProjectStatus
	// basicInfo
	basicInfo := new(BasicInfo)
	basicInfo.ProjectNo = p.SimpleProject.ProjectId
	basicInfo.ProjectName = p.SimpleProject.ProjectName
	basicInfo.Client = p.SimpleProject.ProjectId
	basicInfo.StartDate = p.SimpleProject.StartDate.Format(DATE_FORMAT)
	basicInfo.LimitDate = p.SimpleProject.LimitDate.Format(DATE_FORMAT)
	basicInfo.ReceiveAmount = p.SimpleProject.ReceiveAmount

	// payment
	payment := new(Payment)
	payment.OperatingWorkByTime = p.SimpleProject.GetEstimateOperatingTime()
	payment.OperatingCost = p.SimpleProject.GetEstimateOperatingCost()
	payment.OtherCost = p.SimpleProject.GetOthersCostAmount()

	// members
	members := []ProjectMemberExtOpeTime{}
	for _, pm := range p.ProjectMembers {
		member := new(ProjectMemberExtOpeTime)
		member.ToViewModel(pm, p.Works)
		members = append(members, *member)
	}

	// otherCosts
	otherCosts := []OtherCost{}
	for _, oc := range p.OtherCosts {
		otherCost := new(OtherCost)
		otherCost.Name = oc.Name
		otherCost.Kind = oc.CostKind.KindId
		otherCost.BuyDate = oc.BuyDate.Format(DATE_FORMAT)
		otherCost.Price = oc.Cost
		otherCosts = append(otherCosts, *otherCost)
	}

	// histories
	histories := []ProjectHistory{}
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

type SimpleProjectItem struct {
	ProjectNo           uint    `json:"no"`
	ProjectName         string  `json:"name"`
	ClientNo            uint    `json:"clientNo"`
	ClientName          string  `json:"clientName"`
	Status              int     `json:"status"`
	StartDate           string  `json:"startDate"`
	LimitDate           string  `json:"limitDate"`
	OperatingWorkByTime float32 `json:"operatingWorkByTime"`
	OperatingCost       int64   `json:"operatingCost"`
	OtherCost           int64   `json:"otherCost"`
	ReceiveAmount       int64   `json:"receiveAmount"`
}

func (sp *SimpleProjectItem) ToViewModel(project model.SimpleProject) {
	sp.ProjectNo = project.ProjectId
	sp.ProjectName = project.ProjectName
	sp.ClientNo = project.Customer.CustomerId
	sp.ClientName = project.Customer.CustomerName
	sp.Status = project.ProjectStatus
	sp.StartDate = project.StartDate.Format(DATE_FORMAT)
	if !project.LimitDate.IsZero() {
		sp.LimitDate = project.LimitDate.Format(DATE_FORMAT)
	} else {
		sp.LimitDate = ""
	}
	sp.OperatingWorkByTime = project.GetEstimateOperatingTime()
	sp.OperatingCost = project.GetEstimateOperatingCost()
	sp.OtherCost = project.GetOthersCostAmount()
	sp.ReceiveAmount = project.ReceiveAmount
}

package controller

import (
	"logicApi/src/main/net/baseonlura/slorder/model"
	"logicApi/src/main/net/baseonlura/slorder/viewModel"
)

func GetProjectList() []viewModel.ProjectListItem {
	// TODO DB connect

	viewModel1 := new(viewModel.ProjectListItem)
	viewModel1.ToViewModel(_createProjectStab())

	var resultArray []viewModel.ProjectListItem
	resultArray = append(resultArray, *viewModel1)

	return resultArray
}

func GetProjectItem(projectNo string) viewModel.ProjectItem {
	// TODO DB connect
	vModel := new(viewModel.ProjectItem)
	vModel.ToViewModel(_createProjectStab())
	return *vModel
}

func _createProjectStab() model.Project {
	project := new(model.Project)
	project.ProjectNo = "P-20220301-0001"
	project.ProjectName = "ペット行動管理システム"
	project.Customer = _createCustomerStab()
	project.Status = model.RECEIVED
	project.StartDate = "2022-03-01"
	project.LimitDate = "2022-04-30"
	project.ReceiveAmount = 20000000
	project.EstimateOpeWorkByTime = -1
	project.EstimateOperatingCost = -1
	project.EstimateOtherCost = -1
	project.ProjectMembers = _createProjectMembersStab()
	project.OtherCosts = _createOthersCostsStab()
	project.ProjectHistories = _createHistoriesStab()

	return *project
}
func _createCustomerStab() model.Customer {
	client := new(model.Customer)
	client.CustomerId = "C-20220301-0001"
	client.Name = "ポメラニアン佐藤"
	return *client
}
func _createProjectMembersStab() []model.ProjectMember {

	work1 := new(model.Work)
	work1.WorkDate = "2022-03-01"
	work1.WorkTime = 9
	work2 := new(model.Work)
	work2.WorkDate = "2022-03-02"
	work2.WorkTime = 10
	work3 := new(model.Work)
	work3.WorkDate = "2022-03-03"
	work3.WorkTime = 11
	works := [...]model.Work{*work1, *work2, *work3}

	member1 := new(model.Member)
	member1.MemberId = "M-20220301-0001"
	member2 := new(model.Member)
	member2.MemberId = "M-20220301-0002"
	member3 := new(model.Member)
	member3.MemberId = "M-20220301-0003"
	member4 := new(model.Member)
	member4.MemberId = "M-20220301-0004"

	projectMember1 := new(model.ProjectMember)
	projectMember1.Member = *member1
	projectMember1.UnitCost = 2500
	projectMember1.Works = works[:]
	projectMember2 := new(model.ProjectMember)
	projectMember2.Member = *member2
	projectMember2.UnitCost = 2700
	projectMember2.Works = works[:]
	projectMember3 := new(model.ProjectMember)
	projectMember3.Member = *member3
	projectMember3.UnitCost = 3000
	projectMember3.Works = works[:]
	projectMember4 := new(model.ProjectMember)
	projectMember4.Member = *member4
	projectMember4.UnitCost = 4500
	projectMember4.Works = works[:]

	members := [...]model.ProjectMember{*projectMember1, *projectMember2,
		*projectMember3, *projectMember4}

	return members[:]
}

func _createOthersCostsStab() []model.OtherCost {
	kind1 := new(model.OtherCostKind)
	kind1.KindId = 1
	kind2 := new(model.OtherCostKind)
	kind2.KindId = 2

	cost1 := new(model.OtherCost)
	cost1.Name = "サーバー01"
	cost1.CostKind = *kind1
	cost1.BuyDate = "2022-03-01"
	cost1.Cost = 5000000
	cost2 := new(model.OtherCost)
	cost2.Name = "武田信玄"
	cost2.CostKind = *kind2
	cost2.BuyDate = "2022-03-04"
	cost2.Cost = 400000

	costs := [...]model.OtherCost{*cost1, *cost2}
	return costs[:]
}

func _createHistoriesStab() []model.ProjectHistory {
	lastName := "受取"
	firstName := "太郎"

	history1 := new(model.ProjectHistory)
	history1.LastName = lastName
	history1.FirstName = firstName
	history1.OperationDate = "2022-02-01"
	history1.OperationKind = model.NEW_CREATE
	history2 := new(model.ProjectHistory)
	history2.LastName = lastName
	history2.FirstName = firstName
	history2.OperationDate = "2022-02-20"
	history2.OperationKind = model.NEW_CREATE

	histories := [...]model.ProjectHistory{*history1, *history2}

	return histories[:]
}

package controller

import (
	"logicApi/src/main/net/baseonlura/slorder/model"
	"logicApi/src/main/net/baseonlura/slorder/viewModel"
)

func GetProjectList() []viewModel.ProjectListItem {
	// TODO DB connect
	project1 := new(model.Project)
	project1.ProjectNo = "P-20220301-0001"
	project1.Name = "ペット行動管理システム"
	client := new(model.Customer)
	client.Name = "ポメラニアン佐藤"
	project1.Customer = *client
	project1.Status = model.RECEIVED

	viewModel1 := new(viewModel.ProjectListItem)
	viewModel1.ToViewModel(*project1)

	var resultArray []viewModel.ProjectListItem
	resultArray = append(resultArray, *viewModel1)

	return resultArray
}

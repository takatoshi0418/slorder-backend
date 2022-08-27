package viewModel

import (
	"logicApi/src/main/net/baseonlura/slorder/model"
)

type ProjectListItem struct {
	ProjectNo   uint   `json:"project_no"`
	ProjectName string `json:"project_name"`
	ClientName  string `json:"client_name"`
	Status      string `json:"status"`
	IsEnabled   bool   `json:"isEnabled"`
}

func (pli *ProjectListItem) ToViewModel(p model.SimpleProject) {
	pli.ProjectNo = p.ProjectId
	pli.ProjectName = p.ProjectName
	pli.ClientName = p.Customer.CustomerName
	pli.Status = p.GetStatus().String()
	pli.IsEnabled = p.GetStatus().IsUnfinished()
}

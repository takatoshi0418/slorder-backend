package viewModel

import (
	"logicApi/src/main/net/baseonlura/slorder/model"
)

type ProjectListItem struct {
	ProjectNo   string `json:"project_no"`
	ProjectName string `json:"project_name"`
	ClientName  string `json:"client_name"`
	Status      string `json:"status"`
	IsEnabled   bool   `json:"isEnabled"`
}

func (pli *ProjectListItem) ToViewModel(p model.Project) {
	pli.ProjectNo = p.ProjectNo
	pli.ProjectName = p.Name
	pli.ClientName = p.Customer.Name
	pli.Status = p.Status.String()
	pli.IsEnabled = p.Status.IsUnfinished()
}

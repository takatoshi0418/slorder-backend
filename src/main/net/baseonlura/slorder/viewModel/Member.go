package viewModel

import "logicApi/src/main/net/baseonlura/slorder/model"

type SelectableMemberItem struct {
	MemberID uint   `json:"no"`
	Name     string `json:"name"`
	UnitCost int    `json:"unit"`
}

func (smi *SelectableMemberItem) ToViewModel(m model.Member) {
	smi.MemberID = m.MemberId
	smi.Name = m.GetFullName()
	smi.UnitCost = m.UnitCost
}

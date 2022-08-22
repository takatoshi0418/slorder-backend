package controller

import (
	"logicApi/src/main/net/baseonlura/slorder/model"
	"logicApi/src/main/net/baseonlura/slorder/viewModel"
)

func GetSelectableMembers() []viewModel.SelectableMemberItem {
	var members []viewModel.SelectableMemberItem
	for _, v := range _createMembersStab() {
		vModel := new(viewModel.SelectableMemberItem)
		vModel.ToViewModel(v)
		members = append(members, *vModel)
	}
	return members[:]
}

func _createMembersStab() []model.Member {
	member1 := new(model.Member)
	member1.MemberId = "M-20220301-0001"
	member1.LastName = "浅井"
	member1.FirstName = "長政"
	member1.UnitCost = 2500
	member2 := new(model.Member)
	member2.MemberId = "M-20220301-0002"
	member2.LastName = "織田"
	member2.FirstName = "信長"
	member2.UnitCost = 4000
	member3 := new(model.Member)
	member3.MemberId = "M-20220301-0003"
	member3.LastName = "徳川"
	member3.FirstName = "家康"
	member3.UnitCost = 3500
	member4 := new(model.Member)
	member4.MemberId = "M-20220301-0004"
	member4.LastName = "豊臣"
	member4.FirstName = "秀吉"
	member4.UnitCost = 3500
	member5 := new(model.Member)
	member5.MemberId = "M-20220301-0005"
	member5.LastName = "織田"
	member5.FirstName = "信成"
	member5.UnitCost = 3500

	members := [...]model.Member{*member1, *member2, *member3,
		*member4, *member5}
	return members[:]
}

package controller

import (
	"logicApi/src/main/net/baseonlura/slorder/db"
	"logicApi/src/main/net/baseonlura/slorder/model"
	"logicApi/src/main/net/baseonlura/slorder/viewModel"
)

/**
 * this methods get list of Members from DB after,
 * returns its converted to ViewModel and error interface.
 *
 * returns
 *  * SelectableMemberItem
 *  * errors Interface
 */
func GetSelectableMembers() ([]viewModel.SelectableMemberItem, error) {

	// DB connect
	connection, err := db.GetDBConnection()
	if err != nil {
		return nil, err
	}

	// data from project table
	var members []model.Member
	result := connection.Connection.
		Model(&model.Member{}).
		Find(&members)

	if result.Error != nil {
		return nil, result.Error
	}

	var memberItems []viewModel.SelectableMemberItem
	for _, memberItem := range members {
		vModel := new(viewModel.SelectableMemberItem)
		vModel.ToViewModel(memberItem)
		memberItems = append(memberItems, *vModel)
	}
	return memberItems[:], nil
}

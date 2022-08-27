package controller

import (
	"logicApi/src/main/net/baseonlura/slorder/model"
	"logicApi/src/main/net/baseonlura/slorder/viewModel"

	"logicApi/src/main/net/baseonlura/slorder/db"
)

/**
 * this methods get list of Customers from DB after,
 * returns its converted to ViewModel and error interface.
 *
 * returns
 *  * SelectableClientItem
 *  * errors Interface
 */
func GetSelectableClients() ([]viewModel.SelectableClientItem, error) {

	// DB connect
	connection, err := db.GetDBConnection()
	if err != nil {
		return nil, err
	}

	// data from project table
	var customers []model.Customer
	result := connection.Connection.
		Model(&model.Customer{}).
		Find(&customers)

	if result.Error != nil {
		return nil, result.Error
	}

	var clients []viewModel.SelectableClientItem
	for _, customer := range customers {
		vModel := new(viewModel.SelectableClientItem)
		vModel.ToViewModel(customer)
		clients = append(clients, *vModel)
	}
	return clients[:], nil
}

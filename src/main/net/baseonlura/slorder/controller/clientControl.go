package controller

import (
	"logicApi/src/main/net/baseonlura/slorder/model"
	"logicApi/src/main/net/baseonlura/slorder/viewModel"
)

func GetSelectableClients() []viewModel.SelectableClientItem {
	var clients []viewModel.SelectableClientItem
	for _, v := range _createClientsStab() {
		vModel := new(viewModel.SelectableClientItem)
		vModel.ToViewModel(v)
		clients = append(clients, *vModel)
	}
	return clients[:]

}

func _createClientsStab() []model.Customer {
	client1 := new(model.Customer)
	client1.CustomerId = "C-20220301-0001"
	client1.Name = "ポメラニアン佐藤"
	client2 := new(model.Customer)
	client2.CustomerId = "C-20220301-0002"
	client2.Name = "三日月農協組合"
	client3 := new(model.Customer)
	client3.CustomerId = "C-20220301-0003"
	client3.Name = "山田商店"
	client4 := new(model.Customer)
	client4.CustomerId = "C-20220301-0004"
	client4.Name = "田中文具店"
	client5 := new(model.Customer)
	client5.CustomerId = "C-20220301-0005"
	client5.Name = "スズキ薬局"
	client6 := new(model.Customer)
	client6.CustomerId = "C-20220301-0006"
	client6.Name = "三日月市"

	clients := [...]model.Customer{*client1, *client2, *client3,
		*client4, *client5, *client6}

	return clients[:]
}

package viewModel

import "logicApi/src/main/net/baseonlura/slorder/model"

type SelectableClientItem struct {
	ClientId uint   `json:"id"`
	Name     string `json:"name"`
}

func (sc *SelectableClientItem) ToViewModel(c model.Customer) {
	sc.ClientId = c.CustomerId
	sc.Name = c.CustomerName
}

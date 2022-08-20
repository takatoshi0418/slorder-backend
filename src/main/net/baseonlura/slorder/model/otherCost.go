package model

type OtherCostKind struct {
	KindId string
	Name   string
}

type OtherCost struct {
	CostId   int64
	Name     string
	CostKind OtherCostKind
	BuyDate  string
	Cost     int64
}

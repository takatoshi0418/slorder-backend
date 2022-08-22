package model

type OtherCostKind struct {
	KindId int
	Name   string
}

type OtherCost struct {
	CostId   int64
	Name     string
	CostKind OtherCostKind
	BuyDate  string
	Cost     int64
}

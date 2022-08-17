package model

type Project struct {
	status        ProjectStatus
	no            string
	name          string
	client        Client
	startDate     string
	limitDate     string
	receiveAmount int64
}

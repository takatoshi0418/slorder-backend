package viewModel

import (
	"logicApi/src/main/net/baseonlura/slorder/model"
)

type SimpleWork struct {
	ProjectId uint   `json:"projectId"`
	MemberId  uint   `json:"memberId"`
	WorkDate  string `json:"workDate"`
	WorkTime  int    `json:"workTime"`
}

func (sw *SimpleWork) ToViewModel(work model.Work) {
	sw.ProjectId = work.ProjectId
	sw.MemberId = work.MemberId
	if !work.WorkDate.IsZero() {
		sw.WorkDate = work.WorkDate.Format(DATE_FORMAT)
	} else {
		sw.WorkDate = ""
	}
	sw.WorkTime = work.WorkTime
}

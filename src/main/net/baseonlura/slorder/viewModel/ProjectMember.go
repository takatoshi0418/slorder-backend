package viewModel

import (
	"logicApi/src/main/net/baseonlura/slorder/model"
)

type ProjectMember struct {
	ProjectID  uint       `json:"projectId"`
	MemberID   uint       `json:"memberId"`
	MemberName string     `json:"name"`
	AssignDate string     `json:"assignDate"`
	RejectDate string     `json:"rejectDate"`
	UnitCost   int        `json:"unit"`
	Work       SimpleWork `json:"work"`
}

func (spm *ProjectMember) ToViewModel(pm model.ProjectMember) {
	spm.ProjectID = pm.ProjectId
	spm.MemberID = pm.MemberId
	spm.MemberName = pm.Member.GetFullName()
	spm.AssignDate = pm.AssignDate.Format(DATE_FORMAT)
	if !pm.RejectDate.IsZero() {
		spm.RejectDate = pm.RejectDate.Format(DATE_FORMAT)
	} else {
		spm.RejectDate = ""
	}
	spm.UnitCost = pm.UnitCost
	work := new(SimpleWork)
	work.ToViewModel(pm.Work)
	spm.Work = *work
}

type ProjectMemberExtOpeTime struct {
	MemberID      uint    `json:"value"`
	UnitCost      int     `json:"unit"`
	AssignDate    string  `json:"assignDate"`
	RejectDate    string  `json:"rejectDate"`
	OperatingTime float32 `json:"operatingTime"`
}

func (spm *ProjectMemberExtOpeTime) ToViewModel(pm model.ProjectMember, works []model.Work) {
	spm.MemberID = pm.MemberId
	spm.UnitCost = pm.UnitCost
	spm.AssignDate = pm.AssignDate.Format(DATE_FORMAT)
	if !pm.RejectDate.IsZero() {
		spm.RejectDate = pm.RejectDate.Format(DATE_FORMAT)
	} else {
		spm.RejectDate = ""
	}

	workTime := 0
	for _, w := range works {
		if pm.MemberId == w.MemberId {
			workTime += w.WorkTime
		}
	}
	spm.OperatingTime = float32(workTime / 60)
}

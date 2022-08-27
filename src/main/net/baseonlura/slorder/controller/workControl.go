package controller

import (
	"errors"
	"logicApi/src/main/net/baseonlura/slorder/db"
	"logicApi/src/main/net/baseonlura/slorder/model"
	"logicApi/src/main/net/baseonlura/slorder/viewModel"
	"time"

	"gorm.io/gorm"
)

func GetWorksBelongProject(projectNo uint64, date string) ([]viewModel.SimpleWork, error) {
	// DB connect
	connection, err := db.GetDBConnection()
	if err != nil {
		return nil, err
	}
	dateTime, err := time.Parse(viewModel.DATE_FORMAT, date)
	if err != nil {
		return nil, err
	}

	// data from ProjectMember table
	var projectMembers []model.ProjectMember
	result := connection.Connection.
		Debug().
		Model(&model.ProjectMember{}).
		Where("project_id = ?", projectNo).
		Where("? BETWEEN assign_date AND IFNULL(reject_date, '9999-12-31')", dateTime).
		Find(&projectMembers)

	if result.Error != nil {
		return nil, result.Error
	}

	// data from Work table
	var works []model.Work
	for _, projectMember := range projectMembers {
		result := connection.Connection.
			Debug().
			Model(&model.Work{}).
			Preload("Member").
			Where(&model.Work{ProjectId: projectMember.ProjectId,
				MemberId: projectMember.MemberId}).
			// Where("DATE_FORMAT(work_date, '%Y-%m-%d') = ?", date).
			Find(&works)

		if result.Error != nil {
			return nil, result.Error
		}
	}

	// convert to viewModel
	var viewModels []viewModel.SimpleWork
	for _, work := range works {
		vm := new(viewModel.SimpleWork)
		vm.ToViewModel(work)
		viewModels = append(viewModels, *vm)
	}

	return viewModels, nil
}

func GetWorkByProjectMember(projectMember model.ProjectMember, date string) (model.Work, error) {
	// DB connect
	connection, err := db.GetDBConnection()
	if err != nil {
		return model.Work{}, err
	}

	_, err = time.Parse(viewModel.DATE_FORMAT, date)
	if err != nil {
		return model.Work{}, err
	}

	// data from Work table
	var work model.Work
	result := connection.Connection.
		Debug().
		Model(&model.Work{}).
		Preload("Member").
		Where(&model.Work{ProjectId: projectMember.ProjectId,
			MemberId: projectMember.MemberId}).
		Where("DATE_FORMAT(work_date, '%Y-%m-%d') = ?", date).
		First(&work)

	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return model.Work{}, result.Error
	}

	return work, nil
}

func SetWorksByProjectMember(vModels []viewModel.ProjectMember) error {
	// DB connect
	connection, err := db.GetDBConnection()
	if err != nil {
		return err
	}

	return connection.Connection.Transaction(func(tx *gorm.DB) error {
		for _, vModel := range vModels {
			// Covert ViewModel to Model.
			work := new(model.Work)
			work.ProjectId = vModel.Work.ProjectId
			work.MemberId = vModel.Work.MemberId
			work.WorkDate, err = time.Parse(viewModel.DATE_FORMAT, vModel.Work.WorkDate)
			work.WorkTime = vModel.Work.WorkTime
			if err != nil {
				return err
			}

			// it new or exists judge a Model.
			createFlag := false
			_, err := getWorkByWork(tx, *work)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					createFlag = true
				} else {
					return err
				}
			}

			var result *gorm.DB
			if createFlag {
				result = tx.Debug().Create(&work)
			} else {
				result = tx.Debug().Save(&work)
			}
			if result.Error != nil {
				return result.Error
			}
		}
		return nil
	})
}

func getWorkByWork(con *gorm.DB, w model.Work) (model.Work, error) {
	var work model.Work
	result := con.
		Debug().
		Model(&model.Work{}).
		Where(&model.Work{ProjectId: w.ProjectId,
			MemberId: w.MemberId}).
		Where("DATE_FORMAT(work_date, '%Y-%m-%d') = ?", w.WorkDate).
		First(&work)

	if result.Error != nil {
		return model.Work{}, result.Error
	}
	return work, nil
}

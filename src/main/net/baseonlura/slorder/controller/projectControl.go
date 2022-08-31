package controller

import (
	"errors"
	"logicApi/src/main/net/baseonlura/slorder/model"
	"logicApi/src/main/net/baseonlura/slorder/viewModel"
	"time"

	"logicApi/src/main/net/baseonlura/slorder/db"

	"gorm.io/gorm"
)

/**
 * this methods get list of Projects from DB after,
 * returns its converted to ViewModel and error interface.
 *
 * returns
 *  * ProjectListItem
 *  * errors Interface
 */
func GetProjectList() ([]viewModel.ProjectListItem, error) {
	// DB connect
	connection, err := db.GetDBConnection()
	if err != nil {
		return []viewModel.ProjectListItem{}, err
	}

	// data from project table
	var projects []model.SimpleProject
	result := connection.Connection.
		Model(&model.SimpleProject{}).
		Find(&projects)

	var customers []model.Customer
	CustomerResult := connection.Connection.
		Debug().
		Model(&model.Customer{}).
		Find(&customers)
	for p := range projects {
		for c := range customers {
			if projects[p].CustomerId == customers[c].CustomerId {
				projects[p].Customer = customers[c]
				break
			}
		}
	}

	if result.Error != nil {
		return []viewModel.ProjectListItem{}, result.Error
	}
	if CustomerResult.Error != nil {
		return []viewModel.ProjectListItem{}, CustomerResult.Error
	}

	// convert to viewModel
	var viewModels []viewModel.ProjectListItem
	for _, project := range projects {
		vm := new(viewModel.ProjectListItem)
		vm.ToViewModel(project)
		viewModels = append(viewModels, *vm)
	}

	return viewModels, nil
}

/**
 * this methods find one Project by ProjectNo from DB.
 * returns its converted to ViewModel and error interface.
 *
 * params
 *  * ProjectNo uint64
 *
 * returns
 *  * ProjectListItem
 *  * errors Interface
 */
func GetProjectItem(projectNo uint64) (viewModel.ProjectItem, error) {
	// DB connect
	connection, err := db.GetDBConnection()
	if err != nil {
		return viewModel.ProjectItem{}, err
	}

	// data from project table
	var project model.Project
	result := connection.Connection.
		Debug().
		Model(&model.Project{}).
		Preload("ProjectMembers", getBetweenCondition("assign_date", "reject_date"), time.Now().Truncate(time.Hour*24)).
		Preload("OtherCosts").
		Preload("Works").
		Preload("ProjectHistories").
		First(&project, projectNo)

	if result.Error != nil {
		return viewModel.ProjectItem{}, result.Error
	}

	// convert to viewModel
	vModel := new(viewModel.ProjectItem)
	vModel.ToViewModel(project)
	return *vModel, nil
}

/**
 * this methods find one Simple Project by ProjectNo from DB.
 * returns its converted to ViewModel and error interface.
 *
 * params
 *  * ProjectNo uint64
 *
 * returns
 *  * ProjectListItem
 *  * errors Interface
 */
func GetSimpleProjectItem(projectNo uint64) (viewModel.SimpleProjectItem, error) {
	// DB connect
	connection, err := db.GetDBConnection()
	if err != nil {
		return viewModel.SimpleProjectItem{}, err
	}

	// data from project table
	var project model.SimpleProject
	result := connection.Connection.
		Model(&model.SimpleProject{}).
		Preload("Customer").
		First(&project, projectNo)

	if result.Error != nil {
		return viewModel.SimpleProjectItem{}, result.Error
	}

	// convert to viewModel
	vModel := new(viewModel.SimpleProjectItem)
	vModel.ToViewModel(project)
	return *vModel, nil
}

/**
 * this methods find any Project Members by ProjectNo and target date from DB.
 * returns its converted to ViewModel and error interface.
 *
 * params
 *  * ProjectNo uint64
 *
 * returns
 *  * ProjectMember
 *  * errors Interface
 */
func GetProjectMemberList(projectNo uint64, date string) ([]viewModel.ProjectMember, error) {
	// DB connect
	connection, err := db.GetDBConnection()
	if err != nil {
		return nil, err
	}

	// date string convert to date for Time type.
	// because parameter type check and, prevention SQL injection
	dateTime, err := time.Parse(viewModel.DATE_FORMAT, date)
	if err != nil {
		return nil, err
	}

	// data from projectMembers table
	var projectMembers []model.ProjectMember
	result := connection.Connection.
		Debug().
		Model(&model.ProjectMember{}).
		Preload("Member").
		Where("project_id = ?", projectNo).
		Where(getBetweenCondition("assign_date", "reject_date"), dateTime).
		Find(&projectMembers)

	if result.Error != nil {
		return nil, result.Error
	}

	for i, projectMember := range projectMembers {
		work, err := GetWorkByProjectMember(projectMember, date)
		if err != nil {
			return nil, result.Error
		}
		projectMembers[i].Work = work
		if err != nil {
			return nil, result.Error
		}
	}

	// convert to viewModel
	var vModels []viewModel.ProjectMember
	for _, projectMember := range projectMembers {
		vModel := new(viewModel.ProjectMember)
		vModel.ToViewModel(projectMember)
		vModels = append(vModels, *vModel)
	}

	return vModels, nil
}

/**
 * this methods create one Project.
 * returns its converted to ViewModel and error interface.
 *
 * params
 *  * ProjectNo uint64
 *
 * returns
 *  * ProjectMember
 *  * errors Interface
 */
func GetPureProject() (viewModel.ProjectItem, error) {
	project := new(model.Project)
	nowDate := getNowDate()
	project.SimpleProject.StartDate = nowDate
	vModel := new(viewModel.ProjectItem)
	vModel.ToViewModel(*project)

	return *vModel, nil
}

func GetOtherCostKinds() ([]model.OtherCostKind, error) {
	// DB connect
	connection, err := db.GetDBConnection()
	if err != nil {
		return []model.OtherCostKind{}, err
	}

	otherCostKinds := []model.OtherCostKind{}
	result := connection.Connection.
		Model(&model.OtherCostKind{}).
		Find(&otherCostKinds)

	if result.Error != nil {
		return []model.OtherCostKind{}, result.Error
	}
	return otherCostKinds, nil
}

func SetProjectByProjectItem(vModel viewModel.ProjectItem) error {
	// DB connect
	connection, err := db.GetDBConnection()
	if err != nil {
		return err
	}

	// convert to project from View Model.
	simpleProject := new(model.SimpleProject)
	projectMembers := []model.ProjectMember{}
	otherCosts := []model.OtherCost{}
	// Simple project
	simpleProject.ProjectId = vModel.BasicInfo.ProjectNo
	simpleProject.ProjectName = vModel.BasicInfo.ProjectName
	simpleProject.ProjectStatus = vModel.Status
	simpleProject.StartDate = viewModel.Str2Date(vModel.BasicInfo.StartDate)
	simpleProject.LimitDate = viewModel.Str2Date(vModel.BasicInfo.LimitDate)
	simpleProject.ReceiveAmount = vModel.BasicInfo.ReceiveAmount
	simpleProject.EstimateOperatingTime = &vModel.Payment.OperatingWorkByTime
	simpleProject.EstimateOperatingCost = &vModel.Payment.OperatingCost
	simpleProject.OthersCostAmount = &vModel.Payment.OtherCost
	simpleProject.LastUpdateDate = viewModel.Str2DateTime(vModel.LastUpdateTime)
	simpleProject.CustomerId = vModel.BasicInfo.Client

	for _, member := range vModel.Members {
		projectMember := new(model.ProjectMember)
		projectMember.ProjectId = simpleProject.ProjectId
		projectMember.MemberId = member.MemberID
		projectMember.AssignDate = viewModel.Str2Date(member.AssignDate)
		projectMember.RejectDate = viewModel.Str2Date(member.RejectDate)
		projectMember.UnitCost = member.UnitCost
		projectMembers = append(projectMembers, *projectMember)
	}

	for _, otherCostVModel := range vModel.OtherCosts {
		otherCost := new(model.OtherCost)
		otherCost.CostId = otherCostVModel.No
		otherCost.CostName = otherCostVModel.Name
		otherCost.BuyDate = viewModel.Str2Date(otherCostVModel.BuyDate)
		otherCost.Cost = otherCostVModel.Price
		otherCost.KindId = otherCostVModel.Kind
		otherCosts = append(otherCosts, *otherCost)
	}

	return connection.Connection.Transaction(func(tx *gorm.DB) error {

		// target project find latest from db,
		// If the project is not found, set the create flag to True.
		// If the LastUpdateTime of the project is later than the LastUpdateTime of the target project,
		// an "illegal exclusive access control" error is returned.
		isCreat := false
		var before model.SimpleProject
		result := tx.Debug().
			Model(&model.SimpleProject{}).
			First(&before, simpleProject.ProjectId)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				isCreat = true
			} else {
				return result.Error
			}
		}
		if before.LastUpdateDate.After(simpleProject.LastUpdateDate) {
			return GetSomeoneUpdatedError()
		}

		if isCreat {
			latest := new(model.SimpleProject)
			latestResult := tx.Debug().
				Model(&model.SimpleProject{}).
				Last(&latest)
			if latestResult.Error != nil {
				return latestResult.Error
			}
			simpleProject.ProjectId = latest.ProjectId + 1
			simpleProject.ProjectStatus = model.ESTIMATION.IntKey()
			projectResult := tx.Debug().
				Create(&simpleProject)
			if projectResult.Error != nil {
				return projectResult.Error
			}
		} else {
			projectResult := tx.Debug().Save(&simpleProject)
			if projectResult.Error != nil {
				return projectResult.Error
			}
		}
		projectMemberResult := setProjectMember(tx, isCreat, simpleProject.ProjectId, projectMembers)
		if projectMemberResult != nil {
			return projectMemberResult
		}
		otherCostResult := setOtherCosts(tx, isCreat, simpleProject.ProjectId, otherCosts)
		if otherCostResult != nil {
			return otherCostResult
		}
		historyResult := setProjectHistory(tx, isCreat,
			model.GetProjectStatusByKey(before.ProjectStatus), simpleProject.GetStatus(),
			simpleProject.ProjectId, simpleProject.ProjectName)
		if historyResult != nil {
			return historyResult
		}
		return err
	})
}

func setProjectMember(tx *gorm.DB, isCreat bool, projectNo uint, projectMembers []model.ProjectMember) error {

	if len(projectMembers) == 0 {
		return GetNoneMembersProject()
	}

	if isCreat {
		for i := range projectMembers {
			projectMembers[i].ProjectId = projectNo
			projectMembers[i].AssignDate = getNowDate()
		}
		result := tx.Debug().
			Select("project_id", "member_id", "assign_date", "unit_cost").
			Create(&projectMembers)
		if result.Error != nil {
			return result.Error
		}
		return nil
	}
	before := []model.ProjectMember{}
	findResult := tx.Debug().
		Model(&model.ProjectMember{}).
		Where("project_id = ?", projectNo).
		Where(getBetweenCondition("assign_date", "reject_date"), getNowDate()).
		Find(&before)
	if findResult.Error != nil {
		return findResult.Error
	}

	for _, add := range projectMembers {
		isAdd := true
		for _, member := range before {
			if add.ProjectId == member.ProjectId &&
				add.MemberId == member.MemberId &&
				add.AssignDate.Equal(member.AssignDate) {
				isAdd = false
				break
			}
		}
		if isAdd {
			add.AssignDate = getNowDate()
			result := tx.Debug().
				Select("project_id", "member_id", "assign_date", "unit_cost").
				Create(&add)
			if result.Error != nil {
				return result.Error
			}
		}
	}

	for _, delete := range before {
		isDelete := true
		for _, member := range projectMembers {
			if member.ProjectId == delete.ProjectId &&
				member.MemberId == delete.MemberId &&
				member.AssignDate.Equal(delete.AssignDate) {
				isDelete = false
				break
			}
		}
		if isDelete {
			delete.RejectDate = getNowDate()
			result := tx.Debug().Save(&delete)
			if result.Error != nil {
				return result.Error
			}
		}
	}
	return nil
}

func setOtherCosts(tx *gorm.DB, isCreate bool, projectNo uint, otherCosts []model.OtherCost) error {

	if len(otherCosts) == 0 {
		return nil
	}

	if isCreate {
		for i := range otherCosts {
			otherCosts[i].ProjectId = projectNo
		}
		result := tx.Debug().Create(otherCosts)
		if result.Error != nil {
			return result.Error
		}
		return nil
	}

	before := []model.OtherCost{}
	findResult := tx.Debug().
		Model(&model.OtherCost{}).
		Where("project_id = ?", projectNo).
		Find(&before)
	if findResult.Error != nil {
		return findResult.Error
	}

	for _, add := range otherCosts {
		isAdd := true
		for _, cost := range before {
			if add.CostId == cost.CostId {
				isAdd = false
				break
			}
		}
		if isAdd {
			add.ProjectId = projectNo
			result := tx.Debug().Create(&add)
			if result.Error != nil {
				return result.Error
			}
		}
	}

	for _, delete := range before {
		isDelete := true
		for _, cost := range otherCosts {
			if cost.CostId == delete.CostId {
				isDelete = false
				break
			}
		}
		if isDelete {
			result := tx.Debug().Delete(&delete, delete.CostId)
			if result.Error != nil {
				return result.Error
			}
		}
	}
	return nil
}

func setProjectHistory(tx *gorm.DB, isCreat bool, beforeStatus model.ProjectStatus, afterStatus model.ProjectStatus, projectNo uint, projectName string) error {

	var key int

	switch {
	case isCreat:
		kind := new(model.ProjectOperationKind)
		key = kind.SaveKey()
	// normal sequence
	case beforeStatus.IsEstimationToReceived(afterStatus):
		key = model.RECEIVE.IntKey()
	case beforeStatus.IsReceivedToDelivered(afterStatus):
		key = model.DELIVERED.IntKey()
	case beforeStatus.IsDeliveredToAcceptanced(afterStatus):
		key = model.ACCEPTANCE.IntKey()
	case beforeStatus.IsAcceptancedToPaymented(afterStatus):
		key = model.PAYMENTED.IntKey()
	// lostOrder sequence
	case beforeStatus.IsEstimationToLostorder(afterStatus):
		key = model.LOSTORDER.IntKey()
	case beforeStatus.IsReceivedToLostorder(afterStatus):
		key = model.LOSTORDER.IntKey()
	// cancel sequence
	case beforeStatus.IsPaymentedToAcceptanced(afterStatus):
		key = model.PAYMENTED_CANCEL.IntKey()
	case beforeStatus.IsAcceptancedToDelivered(afterStatus):
		key = model.ACCEPTANCE_CANCEL.IntKey()
	case beforeStatus.IsDeliveredToReceived(afterStatus):
		key = model.DELIVERY_CANCEL.IntKey()
	case beforeStatus.IsReceivedToEstimation(afterStatus):
		key = model.RECEIVE_CANCEL.IntKey()
	case beforeStatus.IsLostorderToReceived(afterStatus):
		key = model.LOST_ORDER_CANCEL.IntKey()
	case beforeStatus.IsLostorderToEstimation(afterStatus):
		key = model.LOST_ORDER_CANCEL.IntKey()
	default:
		kind := new(model.ProjectOperationKind)
		key = kind.SaveKey()
	}

	// get operation user
	opeUser := new(model.User)
	userResult := tx.Debug().
		Model(&model.User{}).
		// TODO get user
		First(&opeUser, 1)
	if userResult.Error != nil {
		return userResult.Error
	}

	history := new(model.ProjectHistory)
	history.ProjectId = projectNo
	history.ProjectName = projectName
	history.UpdateUserId = opeUser.UserId
	history.FirstName = opeUser.FirstName
	history.LastName = opeUser.LastName
	history.OperationDate = getNowDateTime()
	history.OperationKind = key

	result := tx.Debug().Create(&history)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

package controller

import (
	"logicApi/src/main/net/baseonlura/slorder/model"
	"logicApi/src/main/net/baseonlura/slorder/viewModel"
	"time"

	"logicApi/src/main/net/baseonlura/slorder/db"
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
		return nil, err
	}

	// data from project table
	var projects []model.SimpleProject
	result := connection.Connection.
		Model(&model.SimpleProject{}).
		Preload("Customer").
		Find(&projects)

	if result.Error != nil {
		return nil, result.Error
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
		Model(&model.Project{}).
		Preload("Customer").
		Preload("ProjectMembers").
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
		Where("? BETWEEN assign_date and IFNULL(reject_date, '9999-12-31')", dateTime).
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

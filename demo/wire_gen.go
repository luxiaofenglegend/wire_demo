// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package demo

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"io"
	"os"
	"wire_demo/demo/common"
	"wire_demo/demo/dal"
	"wire_demo/demo/handler"
	"wire_demo/demo/service"
)

// Injectors from wire.go:

func InitializeBroadCast() BroadCast {
	message := NewMessage()
	channel := NewChannel(message)
	broadCast := NewBroadCast(channel)
	return broadCast
}

func BuildInjector() (*Injector, error) {
	db := common.InitGormDB()
	projectDal := dal.NewProjectDal(db)
	questionDal := dal.NewQuestionDal(db)
	questionModelDal := dal.NewQuestionModelDal(db)
	projectService := service.NewProjectService(projectDal, projectDal, questionDal, questionModelDal)
	projectHandler := handler.NewProjectHandler(projectService)
	injector := NewInjector(projectHandler)
	return injector, nil
}

func BuildInjectorBySet() (*Injector, error) {
	db := common.InitGormDB()
	projectDal := dal.NewProjectDal(db)
	questionDal := dal.NewQuestionDal(db)
	questionModelDal := dal.NewQuestionModelDal(db)
	projectService := service.NewProjectService(projectDal, projectDal, questionDal, questionModelDal)
	projectHandler := handler.NewProjectHandler(projectService)
	injector := NewInjector(projectHandler)
	return injector, nil
}

func BuildInjectorByStructProvider() (*Injector, error) {
	db := common.InitGormDB()
	projectDal := dal.NewProjectDal(db)
	questionDal := dal.NewQuestionDal(db)
	questionModelDal := dal.NewQuestionModelDal(db)
	projectService := &service.ProjectService{
		IProjectDal:      projectDal,
		ProjectDal:       projectDal,
		QuestionDal:      questionDal,
		QuestionModelDal: questionModelDal,
	}
	projectHandler := handler.NewProjectHandler(projectService)
	injector := NewInjector(projectHandler)
	return injector, nil
}

func BuildInjectorClean() (*Injector, func(), error) {
	db := common.InitGormDB()
	projectDal := dal.NewProjectDal(db)
	questionDal := dal.NewQuestionDal(db)
	questionModelDal := dal.NewQuestionModelDal(db)
	projectService, cleanup, err := service.NewProjectServiceClean(projectDal, projectDal, questionDal, questionModelDal)
	if err != nil {
		return nil, nil, err
	}
	projectHandler := handler.NewProjectHandler(projectService)
	injector := NewInjector(projectHandler)
	return injector, func() {
		cleanup()
	}, nil
}

func BuildByInterfaceValueProvider() *common.MyReader {
	reader := _wireFileValue
	myReader := common.NewMyReader(reader)
	return myReader
}

var (
	_wireFileValue = os.Stdin
)

func BuildByValueProvider() *dal.QuestionDal {
	db := _wireDBValue
	questionDal := dal.NewQuestionDal(db)
	return questionDal
}

var (
	_wireDBValue = &gorm.DB{}
)

// wire.go:

var ProjectSet = wire.NewSet(handler.NewProjectHandler, service.NewProjectService, wire.Bind(new(dal.IProjectDal), new(*dal.ProjectDal)), dal.NewProjectDal, dal.NewQuestionDal, dal.NewQuestionModelDal)

var ISet = wire.NewSet(wire.InterfaceValue(new(io.Reader), os.Stdin))
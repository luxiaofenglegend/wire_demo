//go:build wireinject

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

func InitializeBroadCast() BroadCast {
	wire.Build(NewBroadCast, NewChannel, NewMessage)
	return BroadCast{}
}

func BuildInjector() (*Injector, error) {
	wire.Build(
		NewInjector,
		handler.NewProjectHandler,
		service.NewProjectService,
		wire.Bind(new(dal.IProjectDal), new(*dal.ProjectDal)),
		dal.NewProjectDal,
		dal.NewQuestionDal,
		dal.NewQuestionModelDal,
		common.InitGormDB,
	)
	return new(Injector), nil
}

func BuildInjectorBySet() (*Injector, error) {
	wire.Build(
		NewInjector,
		ProjectSet,
		common.InitGormDB,
	)
	return new(Injector), nil
}

var ProjectSet = wire.NewSet(
	handler.NewProjectHandler,
	service.NewProjectService,
	wire.Bind(new(dal.IProjectDal), new(*dal.ProjectDal)),
	dal.NewProjectDal,
	dal.NewQuestionDal,
	dal.NewQuestionModelDal,
)

func BuildInjectorByStructProvider() (*Injector, error) {
	wire.Build(
		NewInjector,
		handler.NewProjectHandler,
		wire.Struct(new(service.ProjectService), "*"),
		// interface bind
		wire.Bind(new(dal.IProjectDal), new(*dal.ProjectDal)),
		dal.NewProjectDal,
		dal.NewQuestionDal,
		dal.NewQuestionModelDal,
		common.InitGormDB,
	)
	return new(Injector), nil
}

func BuildInjectorClean() (*Injector, func(), error) {
	wire.Build(
		NewInjector,
		handler.NewProjectHandler,
		service.NewProjectServiceClean,
		wire.Bind(new(dal.IProjectDal), new(*dal.ProjectDal)),
		dal.NewProjectDal,
		dal.NewQuestionDal,
		dal.NewQuestionModelDal,
		common.InitGormDB,
	)
	return new(Injector), nil, nil
}

var ISet = wire.NewSet(wire.InterfaceValue(new(io.Reader), os.Stdin))

func BuildByInterfaceValueProvider() *common.MyReader {
	wire.Build(common.NewMyReader, ISet)
	return new(common.MyReader)
}

func BuildByValueProvider() *dal.QuestionDal {
	wire.Build(
		dal.NewQuestionDal,
		wire.Value(&gorm.DB{}),
	)
	return new(dal.QuestionDal)
}

package service

import (
	"context"
	"wire_demo/demo/bo"
	"wire_demo/demo/dal"
)

type ProjectService struct {
	IProjectDal      dal.IProjectDal
	ProjectDal       *dal.ProjectDal
	QuestionDal      *dal.QuestionDal
	QuestionModelDal *dal.QuestionModelDal
}

func NewProjectService(iProjectDal dal.IProjectDal, projectDal *dal.ProjectDal, questionDal *dal.QuestionDal, questionModelDal *dal.QuestionModelDal) *ProjectService {
	return &ProjectService{
		IProjectDal:      iProjectDal,
		ProjectDal:       projectDal,
		QuestionDal:      questionDal,
		QuestionModelDal: questionModelDal,
	}
}

func NewProjectServiceClean(iProjectDal dal.IProjectDal, projectDal *dal.ProjectDal, questionDal *dal.QuestionDal, questionModelDal *dal.QuestionModelDal) (*ProjectService, func(), error) {
	clean := func() {

	}
	return &ProjectService{
		IProjectDal:      iProjectDal,
		ProjectDal:       projectDal,
		QuestionDal:      questionDal,
		QuestionModelDal: questionModelDal,
	}, clean, nil
}

func (s *ProjectService) Create(ctx context.Context, projectBo *bo.ProjectCreateBo) (int64, error) {
	return 0, nil
}

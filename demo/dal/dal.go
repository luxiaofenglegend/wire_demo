package dal

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"wire_demo/demo/entity"
)

type IProjectDal interface {
	Create(ctx context.Context, item *entity.Project) (err error)
	// ...
}

type ProjectDal struct {
	DB *gorm.DB
}

func NewProjectDal(db *gorm.DB) *ProjectDal {
	return &ProjectDal{
		DB: db,
	}
}

func (dal *ProjectDal) Create(ctx context.Context, item *entity.Project) error {
	result := dal.DB.Create(item)
	return errors.WithStack(result.Error)
}

// QuestionDal、QuestionModelDal...

type QuestionDal struct {
	DB *gorm.DB
}

func NewQuestionDal(db *gorm.DB) *QuestionDal {
	return &QuestionDal{
		DB: db,
	}
}

type QuestionModelDal struct {
	DB *gorm.DB
}

func NewQuestionModelDal(db *gorm.DB) *QuestionModelDal {
	return &QuestionModelDal{
		DB: db,
	}
}

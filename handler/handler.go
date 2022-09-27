package handler

import (
	"context"
	"wire_demo/demo/project"
	"wire_demo/demo/service"
)

type ProjectHandler struct {
	ProjectService *service.ProjectService
}

func NewProjectHandler(srv *service.ProjectService) *ProjectHandler {
	return &ProjectHandler{
		ProjectService: srv,
	}
}

func (s *ProjectHandler) CreateProject(ctx context.Context, req *project.CreateProjectRequest) (resp *project.CreateProjectResponse, err error) {
	return nil, nil
}

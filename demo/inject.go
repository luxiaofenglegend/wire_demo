package demo

import (
	"wire_demo/demo/handler"
)

type Injector struct {
	ProjectHandler *handler.ProjectHandler
	// components，others...
}

func NewInjector(handler *handler.ProjectHandler) *Injector {
	return &Injector{
		ProjectHandler: handler,
	}
}

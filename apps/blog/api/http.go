package api

import (
	"github.com/ACK-lcn/Blog/apps/blog"
	"github.com/ACK-lcn/Blog/ioc"
)

func Init() {
	ioc.ApiHandler().Register(&apiHandler{})
}

type apiHandler struct {
	svc blog.Service
}

func (t *apiHandler) Init() error {
	t.svc = ioc.Controller().Get(blog.AppName).(blog.Service)
	return nil
}

func (t *apiHandler) Name() string {
	return blog.AppName
}

func NewTokenApiHandler(tokenServiceImpl blog.Service) *apiHandler {
	return &apiHandler{
		svc: tokenServiceImpl,
	}
}

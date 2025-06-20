package ioc

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type IocContainter struct {
	store map[string]iocObject
}

// Init all objects in the container
func (c *IocContainter) Init() error {
	for _, obj := range c.store {
		if err := obj.Init(); err != nil {
			return err
		}
	}
	return nil
}

func (c *IocContainter) Register(obj iocObject) {
	c.store[obj.Name()] = obj
}

func (c *IocContainter) Get(name string) any {
	return c.store[name]
}

type GinApiHandler interface {
	Registry(r gin.IRouter)
}

type GrpcHandler interface {
	Registry(r *grpc.Server)
}

func (c *IocContainter) RouteRegistry(r gin.IRouter) {
	for _, obj := range c.store {
		if api, ok := obj.(GinApiHandler); ok {
			api.Registry(r)
		}
	}
}

func (c *IocContainter) GrpcServerRegistry(r *grpc.Server) {
	for _, obj := range c.store {
		if grpcServerImpl, ok := obj.(GrpcHandler); ok {
			grpcServerImpl.Registry(r)
		}
	}
}

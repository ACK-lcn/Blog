package protocol

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ACK-lcn/Blog/conf"
	"github.com/ACK-lcn/Blog/ioc"
	"github.com/gin-gonic/gin"
)

func NewHttpServer() *HttpServer{
	r := gin.Default()
	ioc.ApiHandler().RouteRegistry(r.Group("/api/blog"))

	return &HttpServer{
		server: &http.Server{
			Addr: conf.C().App.HttpAddress(),
			Handler: r,
		},
	}
}


type HttpServer struct{
	server *http.Server
}


func (s *HttpServer) Run() error{
	fmt.Printf("listen addr: %s\n", conf.C().App.HttpAddress())
	return  s.server.ListenAndServe()
}

// shuntdown
func (s *HttpServer) Close(ctx context.Context){
	s.server.Shutdown(ctx)
}
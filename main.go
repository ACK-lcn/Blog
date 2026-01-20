package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	// _ "github.com/ACK-lcn/Blog/apps"

	"github.com/ACK-lcn/Blog/apps/protocol"
	"github.com/ACK-lcn/Blog/conf"
	"github.com/ACK-lcn/Blog/ioc"
	"google.golang.org/grpc"
)

func main() {
	// Load configuration
	err := conf.LoadConfigFromFile("etc/config.toml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Initialize controller
	// userServiceImpl := userImpl.NewUserServiceImpl()
	// tokenServiceImpl := tokenImpl.NewTokenServiceImpl(userServiceImpl)
	// tokenApiHandler := tokenApiHandler.NewTokenApiHandler(tokenServiceImpl)

	// // Start the server，Register handler route.
	// r := gin.Default()
	// tokenApiHandler.Register(r.Group("/api/blog"))
	// addr := conf.C().App.HttpAddress()
	// r.Run(addr)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	if err := ioc.Controller().Init(); err != nil{
		fmt.Println(err)
		os.Exit(1)
	}

	// Init ApiHeadler
	if err := ioc.ApiHandler().Init(); err != nil{
		fmt.Println(err)
		os.Exit(1)
	}

	httpServer := protocol.NewHttpServer()
	go func(){
		if err := httpServer.Run(); err != nil{
			fmt.Printf("start http server error, %s\n", err)
		}
	}()

	grpcServer := grpc.NewServer()
	ioc.Controller().GrpcServerRegistry(grpcServer)

	go func(){
		fmt.Println("grpc 服务访问地址: 127.0.0.1:8899")
		lis, err := net.Listen("tcp", ":8899")
		if err != nil{
			log.Fatal(err)
		}

		err = grpcServer.Serve(lis)
		if err != nil{
			fmt.Printf("start grpc server eror, %s", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)
	sn := <-ch
	fmt.Println(sn)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	httpServer.Close(ctx)
}
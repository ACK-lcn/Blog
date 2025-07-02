package main

import (
	"fmt"
	"os"

	// _ "github.com/ACK-lcn/Blog/apps"
	tokenApiHandler "github.com/ACK-lcn/Blog/apps/token/api"
	tokenImpl "github.com/ACK-lcn/Blog/apps/token/impl"
	userImpl "github.com/ACK-lcn/Blog/apps/user/impl"
	"github.com/ACK-lcn/Blog/conf"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	err := conf.LoadConfigFromFile("etc/config.toml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Initialize controller
	userServiceImpl := userImpl.NewUserServiceImpl()
	tokenServiceImpl := tokenImpl.NewTokenServiceImpl(userServiceImpl)
	tokenApiHandler := tokenApiHandler.NewTokenApiHandler(tokenServiceImpl)

	// Start the server，Register handler route.
	r := gin.Default()
	tokenApiHandler.Register(r.Group("/api/blog"))
	addr := conf.C().App.HttpAddress()
	r.Run(addr)
	if err != nil {
		fmt.Println(err)
	}
}

package test

import (
	"github.com/ACK-lcn/Blog/conf"
	"github.com/ACK-lcn/Blog/ioc"
)

func DevelopmentSetup() {
	err := conf.LoadConfigFromEnv()
	if err != nil {
		panic(err)
	}

	if err := ioc.Controller().Init(); err != nil {
		panic(err)
	}
}

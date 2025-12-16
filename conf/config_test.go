package conf_test

import (
	"testing"

	"github.com/ACK-lcn/Blog/conf"
)

func TestLoadConfigFromToml(t *testing.T) {
	err := conf.LoadConfigFromFile("test/config.toml")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(conf.C())
}

func TestLoadConfigFromEnv(t *testing.T) {
	err := conf.LoadConfigFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(conf.C())
}

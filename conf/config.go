package conf

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DefaultConfig() *Config {
	return &Config{
		MySQL: &MySQL{
			Host:     "0.0.0.0",
			DB:       "blog",
			Port:     3306,
			Username: "blog",
			Password: "blog123",
		},
		App: &App{
			HttpHost: "0.0.0.0",
			HttpPort: 9080,
		},
	}
}

type Config struct {
	MySQL *MySQL `json:"mysql" toml:"mysql"`
	App   *App   `json:"app" toml:"app"`
}

func (c *Config) String() string {
	dj, err := json.Marshal(c)
	if err != nil {
		fmt.Println("marshal config error")
	}
	return string(dj)
}

type MySQL struct {
	Host     string `json:"host" toml:"host" env:"MYSQL_HOST"`
	Port     int    `json:"port" toml:"port" env:"MYSQL_PORT"`
	DB       string `json:"db" toml:"db" env:"MYSQL_DB"`
	Username string `json:"username" toml:"username" env:"MYSQL_USERNAME"`
	Password string `json:"password" toml:"password" env:"MYSQL_PASSWORD"`

	// cache object
	lock sync.Mutex
	conn *gorm.DB
}

// eg: https://gorm.io/zh_CN/docs/connecting_to_the_database.html
func (m *MySQL) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.DB,
	)
}

func (m *MySQL) GetConnection() *gorm.DB {
	m.lock.Lock()
	defer m.lock.Unlock()

	// if not, assign a value
	if m.conn == nil {
		conn, err := gorm.Open(mysql.Open(m.DSN()), &gorm.Config{})
		if err != nil {
			log.Panicln("mysql connection error")
		}
		m.conn = conn
	}
	return m.conn
}

type App struct {
	HttpHost string `json:"http_host" env:"HTTP_HOST"`
	HttpPort int64  `json:"http_port" env:"HTTP_PORT"`
}

func (a *App) HttpAddress() string {
	return fmt.Sprintf("%s:%d", a.HttpHost, a.HttpPort)
}

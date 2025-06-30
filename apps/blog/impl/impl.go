package impl

import (
	"github.com/ACK-lcn/Blog/apps/blog"
	"github.com/ACK-lcn/Blog/conf"
	"github.com/ACK-lcn/Blog/ioc"
	"gorm.io/gorm"
)

func init() {
	ioc.Controller().Register(&blogServiceImpl{})
}

type blogServiceImpl struct {
	// db
	db *gorm.DB
}

// Init implements ioc.iocObject.
func (i *blogServiceImpl) Init() error {
	i.db = conf.C().MySQL.GetConnection().Debug()
	return nil
}

// Name implements ioc.iocObject.
func (i *blogServiceImpl) Name() string {
	return blog.AppName
}

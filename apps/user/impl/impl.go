package impl

import (
	"context"

	"github.com/ACK-lcn/Blog/apps/user"
	"github.com/ACK-lcn/Blog/conf"
	"github.com/ACK-lcn/Blog/exception"
	"github.com/ACK-lcn/Blog/ioc"
	"gorm.io/gorm"
)

// Register UserServiceImpl to Ioc.
func Init() {
	ioc.Controller().Register(&UserServiceImpl{})
}

// Used to explicitly constrain the implementation of the interface
var _ user.Service = &UserServiceImpl{}

// Both need to be called after initial configuration.
func NewUserServiceImpl() *UserServiceImpl {
	return &UserServiceImpl{
		db: conf.C().MySQL.GetConnection(),
		// Turn on debug
		// db: conf.C().MySQL.GetConnection().Debug(),
	}
}

type UserServiceImpl struct {
	db *gorm.DB
}

// Init implements ioc.iocObject.
func (i *UserServiceImpl) Init() error {
	i.db = conf.C().MySQL.GetConnection().Debug()
	return nil
}

// Name implements ioc.iocObject.
func (i *UserServiceImpl) Name() string {
	return user.AppName
}

// Create User
func (i *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.User, error) {
	// Verify user parameters
	if err := req.Validata(); err != nil {
		return nil, err
	}

	// Generater user object
	ins := user.NewUser(req)

	// save db
	/* When the context has a cancel operation, it will be monitored that
	"withContext has its own implementation of a similar Context down monitor";
	This mechanism will not continue to operate on the database (cancel the database stuck operation).
	*/
	if err := i.db.WithContext(ctx).Create(ins).Error; err != nil {
		return nil, err
	}

	// return
	return ins, nil
}

// Describe User
func (i *UserServiceImpl) DescribeUser(ctx context.Context, req *user.DescribeUserRequest) (*user.User, error) {
	query := i.db.WithContext(ctx)

	switch req.DescribeBy {
	case user.DESCRIBE_BY_ID:
		query = query.Where("id=?", req.DescribeValue)
	case user.DESCRIBE_BY_USERNAME:
		query = query.Where("username=?", req.DescribeValue)
	}

	ins := user.NewUser(user.NewCreateUserRequest())
	if err := query.First(ins).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.NewNotFound("user %s not found", req.DescribeValue)
		}
		return nil, err
	}

	// The hash value is stored in the database.
	ins.SetIsHashed()

	return ins, nil
}

// Delete User
func (i *UserServiceImpl) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) error {
	_, err := i.DescribeUser(ctx, user.NewDescribeUserRequestById(req.IdString()))
	if err != nil {
		return err
	}

	return i.db.WithContext(ctx).Where("id=?", req.Id).Delete(&user.User{}).Error
}

package apps

// Business controller
// (responsible for importing all business implementations: registered in the Controller area in Ioc)
import (
	// Note: The order in which packages are imported is the order in which objects are registered.
	_ "github.com/ACK-lcn/Blog/apps/token/impl"
	_ "github.com/ACK-lcn/Blog/apps/user/impl"
	// Api Handler注册
)

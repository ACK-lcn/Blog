package ioc

func Controller() *IocContainter {
	return controllerContainer
}

var controllerContainer = &IocContainter{
	store: map[string]iocObject{},
}

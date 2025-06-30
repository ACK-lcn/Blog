package ioc

// Used to register Controller objects.
func ApiHandler() *IocContainter {
	return apiHandlerContainer
}

// ioc registry object, globally only.
var apiHandlerContainer = &IocContainter{
	store: map[string]iocObject{},
}

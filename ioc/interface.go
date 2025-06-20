package ioc

// Define constraints for registered objects.
type iocObject interface {
	// Object Init
	Init() error
	// Object Name
	Name() string
}

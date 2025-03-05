package databaseDrivers

type Driver interface {
	DriverName() string
	BuildNamedArgs(args ...map[string]any) any
}

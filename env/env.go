package env

import (
	"os"
	"runtime/debug"
	"slices"
)

var trueVals = []string{"1", "True", "TRUE", "true"}

func getenv(key, dflt string) func() string {
	return func() string {
		if v := os.Getenv(key); v != "" {
			return v
		}
		return dflt
	}
}

var APP_HOST = getenv("APP_HOST", "0.0.0.0")
var APP_PORT = getenv("APP_PORT", "8000")
var APP_NAME = getenv("APP_NAME", "")
var ENV = getenv("ENV", "")
var LOG_LEVEL = getenv("LOG_LEVEL", "debug")

func getBoolEnv(key, dflt string) func() bool {
	return func() bool { return slices.Contains(trueVals, getenv(key, dflt)()) }
}

var DEBUG = getBoolEnv("DEBUG", "False")

var VERSION = func() string {
	info, ok := debug.ReadBuildInfo()
	if !ok || info == nil {
		return "0.0.0"
	}
	version := info.Main.Version
	if version != "" {
		return version
	}
	return "0.0.0"
}

var GOVERSION = func() string {
	info, ok := debug.ReadBuildInfo()
	if !ok || info == nil {
		return "0.0.0"
	}
	version := info.GoVersion
	if version != "" {
		return version
	}
	return "0.0.0"
}

var BASE_PATH = func() string {
	info, ok := debug.ReadBuildInfo()
	if !ok || info == nil {
		return ""
	}
	return info.Main.Path
}

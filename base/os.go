package base

import "os"

// AppOs holds the handle to the current implementation of the `Os` interface.
// In production, this would be set to the `DefaultOs` implementation.
// In testing, this would be manipulated using mocks.
var AppOs Os = &DefaultOs{}

// Os interface encapsulates the Go's standard library functions for the `os` package.
// Below contract is not exhaustive. Instead, they contain methods currently
// being used in the production code. As more of the `os` package are being used in
// production, those methods must make their way into this interface so that it can be
// mocked accordingly.
// The below comment is important. It's a directive to the `mockgen` tool to generate mocks
// when run using `go generate ./...` in command line.
//
//go:generate mockgen -destination=../mocks/mock_os.go -package=mocks github.com/kart/go-mocking-tutorial/base Os
type Os interface {
	// Environment API
	Getenv(key string) string
	LookupEnv(key string) (string, bool)
	// Network API
	Hostname() (string, error)
	// File APIs
	Create(name string) (File, error)
	Open(name string) (File, error)
	OpenFile(name string, flag int, perm os.FileMode) (File, error)
	Stat(name string) (os.FileInfo, error)
	Chdir(name string) error
	Remove(name string) error
	Getwd() (string, error)
	// Process APIs
	Getpid() int
	Executable() (string, error)
	// System APIs
	Exit(code int)
}

// DefaultOs is the default implementation of the `Os` interface.
type DefaultOs struct{}

// Getenv is the default implementation of `os.Getenv`
func (DefaultOs) Getenv(key string) string {
	return os.Getenv(key)
}

// LookupEnv is the default implementation of `os.LookupEnv`
func (DefaultOs) LookupEnv(key string) (string, bool) {
	return os.LookupEnv(key)
}

// Hostname is the default implementation of `os.Hostname`
func (DefaultOs) Hostname() (string, error) {
	return os.Hostname()
}

// Create is the default implementation of `os.Create`
func (DefaultOs) Create(name string) (File, error) {
	return os.Create(name)
}

// Open is the default implementation of `os.Open`
func (DefaultOs) Open(name string) (File, error) {
	return os.Open(name)
}

// OpenFile is the default implementation of `os.OpenFile`
func (DefaultOs) OpenFile(name string, flag int, perm os.FileMode) (File, error) {
	return os.OpenFile(name, flag, perm)
}

// Stat ...
func (DefaultOs) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

// Chdir ...
func (DefaultOs) Chdir(name string) error {
	return os.Chdir(name)
}

// Remove ...
func (DefaultOs) Remove(name string) error {
	return os.Remove(name)
}

// Getwd ...
func (DefaultOs) Getwd() (string, error) {
	return os.Getwd()
}

// Getpid ...
func (DefaultOs) Getpid() int {
	return os.Getpid()
}

// Executable ...
func (DefaultOs) Executable() (string, error) {
	return os.Executable()
}

// Exit ...
func (DefaultOs) Exit(code int) {
	os.Exit(code)
}

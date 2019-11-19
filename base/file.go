package base

import (
	"io"
	"os"
)

// File interface encapsulates the Go's standard library `os.File` structure.
// Below contract is not exhaustive. Instead, they contain methods currently used
// in the production code. As more of the `os.File` APIs are being used in
// production, those methods must make its way into this interface so that they
// can be properly mocked.
// The below comment is important. It's a directive to the `mockgen` tool to generate mocks
//go:generate mockgen -destination=../mocks/mock_file.go -package=mocks github.com/kart/go-mocking-tutorial/base File
type File interface {
	// I/O operations
	io.Closer
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Writer
	io.WriterAt

	// Return the name of the file.
	Name() string
	// Meta operations
	Chdir() error
	Chmod(mode os.FileMode) error
	Chown(uid, gid int) error
	Stat() (os.FileInfo, error)
	// Derived I/O operations
	WriteString(s string) (n int, err error)
	Sync() error
	Truncate(size int64) error
	// Directory operations.
	Readdir(count int) ([]os.FileInfo, error)
	Readdirnames(n int) (names []string, err error)
}

package filesystem

import (
	"errors"
	"io"
)

var (
	ErrFileNotFound                   = errors.New("file not found")
	ErrFileFound                      = errors.New("file already exists")
	ErrDirectoryFound                 = errors.New("directory already exists")
	ErrUnresolvableDirectoryStructure = errors.New("unresolvable directory structure")
	ErrDirectory                      = errors.New("location contains directory but handler expects file")
	ErrFile                           = errors.New("location contains file but handler expects directory")
	ErrWriteLengthMismatch            = errors.New("amount of written bytes does not match the reported amount")
	ErrUnsupportedContentOperation    = errors.New("content operation is not supported")
)

type Filesystem interface {
	handleReadContentOf(string) (Content, error)
	handleStreamContentOf(string) (io.ReadCloser, error)
	handleCheckIfExists(string) (bool, error)
	handleCreateFile(string, ...Argument) error
	handleWriteContentTo(string, []byte, ...Argument) error
	handleStreamContentTo(string, io.Reader, ...Argument) error
	handleCreateDirectory(string, ...Argument) error
}

package filesystem

import (
	"errors"
	"io"
)

var ErrFileNotFound = errors.New("file not found")

type Filesystem interface {
	handleReadContentOf(string) (Content, error)
	handleStreamContentOf(string) (io.ReadCloser, error)
	handleCheckIfExists(string) (bool, error)
}

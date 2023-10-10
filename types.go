package filesystem

import (
	"errors"
	"io"
)

var ErrFileNotFound = errors.New("file not found")

type (
	Streamer interface {
		io.Reader
		io.Closer
	}
	Filesystem interface {
		handleReadContentOf(string) (Content, error)
		handleStreamContentOf(string) (Streamer, error)
	}
)

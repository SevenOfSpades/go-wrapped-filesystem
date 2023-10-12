package filesystem

import (
	"io"
	"os"
)

type (
	// ReadContentOfHandlerFunc is expected to be provided for as handler for ReadContentOf.
	ReadContentOfHandlerFunc func(string) (Content, error)

	// StreamContentOfHandlerFunc is expected to be provided for as handler for StreamContentOf.
	StreamContentOfHandlerFunc func(string) (io.ReadCloser, error)

	// CheckIfExistsHandlerFunc is expected to be provided for as handler for CheckIfExistsHandlerFunc.
	CheckIfExistsHandlerFunc func(string) (bool, error)
)

func readContentOfDefaultHandler(path string) (Content, error) {
	s, err := streamContentOfDefaultHandler(path)
	if err != nil {
		return nil, err
	}
	res, err := io.ReadAll(s)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func streamContentOfDefaultHandler(path string) (io.ReadCloser, error) {
	f, err := os.OpenFile(path, os.O_RDONLY, os.FileMode(0600)) //nolint:gosec
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrFileNotFound
		}
		return nil, err
	}
	return f, nil
}

func checkIfExistsDefaultHandler(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

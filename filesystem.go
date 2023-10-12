package filesystem

import (
	"fmt"
	"io"

	"github.com/SevenOfSpades/go-just-options"
)

// New will create new Filesystem instance.
// Depending on use case all or some handlers for dedicated functions can be overridden with options.Option.
// By default, Filesystem will use builtin functions for interacting with filesystem.
func New(opts ...options.Option) (Filesystem, error) {
	opt := options.New().Resolve(opts...)

	optReadContentOfHandler, err := options.ReadOrDefault[ReadContentOfHandlerFunc](opt, optionReadContentOfHandler, readContentOfDefaultHandler)
	if err != nil {
		return nil, fmt.Errorf("filesystem initialization failed: %w", err)
	}
	optStreamContentOfHandler, err := options.ReadOrDefault[StreamContentOfHandlerFunc](opt, optionStreamContentOfHandler, streamContentOfDefaultHandler)
	if err != nil {
		return nil, fmt.Errorf("filesystem initialization failed: %w", err)
	}
	optCheckIfExistsHandler, err := options.ReadOrDefault[CheckIfExistsHandlerFunc](opt, optionCheckIfExistsHandler, checkIfExistsDefaultHandler)
	if err != nil {
		return nil, fmt.Errorf("filesystem initialization failed: %w", err)
	}

	return newFilesystem(optReadContentOfHandler, optStreamContentOfHandler, optCheckIfExistsHandler)
}

// ReadContentOf will return entire content of file from provided path.
// If file does not exist it will return ErrFileNotFound error.
func ReadContentOf(fs Filesystem, path string) (Content, error) {
	res, err := fs.handleReadContentOf(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read content of %s: %w", path, err)
	}
	return res, nil
}

// StreamContentOf will return pointer to struct implementing io.ReadCloser interface.
// Implementation for this specific handler may or may not stream content directly from it source.
// Some packages or sources for files require preloading entire content before performing any operation on it.
//
// If file does not exist it will return ErrFileNotFound error.
func StreamContentOf(fs Filesystem, path string) (io.ReadCloser, error) {
	res, err := fs.handleStreamContentOf(path)
	if err != nil {
		return nil, fmt.Errorf("failed to attach reader to %s: %w", path, err)
	}
	return res, nil
}

// CheckIfExists will verify if file/directory exists on provided path.
func CheckIfExists(fs Filesystem, path string) (bool, error) {
	res, err := fs.handleCheckIfExists(path)
	if err != nil {
		return false, fmt.Errorf("failed to verify existence of %s: %w", path, err)
	}
	return res, nil
}

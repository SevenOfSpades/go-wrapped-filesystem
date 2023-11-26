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
	opt := options.Resolve(opts)

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
	optCreateFileHandler, err := options.ReadOrDefault[CreateFileHandlerFunc](opt, optionCreateFileHandler, createFileDefaultHandler)
	if err != nil {
		return nil, fmt.Errorf("filesystem initialization failed: %w", err)
	}
	optWriteContentToHandler, err := options.ReadOrDefault[WriteContentToHandlerFunc](opt, optionWriteContentToHandler, writeContentToDefaultHandler)
	if err != nil {
		return nil, fmt.Errorf("filesystem initialization failed: %w", err)
	}
	optStreamContentToHandler, err := options.ReadOrDefault[StreamContentToHandlerFunc](opt, optionStreamContentToHandler, streamContentToDefaultHandler)
	if err != nil {
		return nil, fmt.Errorf("filesystem initialization failed: %w", err)
	}
	optCreateDirectoryHandler, err := options.ReadOrDefault[CreateDirectoryHandlerFunc](opt, optionCreateDirectoryHandler, createDirectoryDefaultHandler)
	if err != nil {
		return nil, fmt.Errorf("filesystem initialization failed: %w", err)
	}

	return newFilesystem(
		optReadContentOfHandler,
		optStreamContentOfHandler,
		optCheckIfExistsHandler,
		optCreateFileHandler,
		optWriteContentToHandler,
		optStreamContentToHandler,
		optCreateDirectoryHandler,
	)
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

// CreateFile creates empty file at provided location (along with missing parts of directory tree if allowed).
func CreateFile(fs Filesystem, path string, args ...Argument) error {
	if err := fs.handleCreateFile(path, args...); err != nil {
		return fmt.Errorf("failed to create file %s: %w", path, err)
	}
	return nil
}

// WriteContentTo appends/overwrites content of file with provided data.
func WriteContentTo[T ~string | ~[]byte](fs Filesystem, path string, content T, args ...Argument) error {
	if err := fs.handleWriteContentTo(path, []byte(content), args...); err != nil {
		return fmt.Errorf("failed to write content to %s: %w", path, err)
	}
	return nil
}

// StreamContentTo appends/overwrites content of file with content of provided io.Reader.
func StreamContentTo(fs Filesystem, path string, content io.Reader, args ...Argument) error {
	if err := fs.handleStreamContentTo(path, content, args...); err != nil {
		return fmt.Errorf("failed to stream content to %s: %w", path, err)
	}
	return nil
}

// CreateDirectory makes empty directory at provided location (along with missing parts of directory tree if allowed).
func CreateDirectory(fs Filesystem, path string, args ...Argument) error {
	if err := fs.handleCreateDirectory(path, args...); err != nil {
		return fmt.Errorf("failed to create directory at %s: %w", path, err)
	}
	return nil
}

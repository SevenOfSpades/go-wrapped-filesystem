package filesystem

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type (
	// ReadContentOfHandlerFunc is expected to be provided for as handler for ReadContentOf.
	ReadContentOfHandlerFunc func(string) (Content, error)

	// StreamContentOfHandlerFunc is expected to be provided for as handler for StreamContentOf.
	StreamContentOfHandlerFunc func(string) (io.ReadCloser, error)

	// CheckIfExistsHandlerFunc is expected to be provided for as handler for CheckIfExists.
	CheckIfExistsHandlerFunc func(string) (bool, error)

	// CreateFileHandlerFunc is expected to be provided for as handler for CreateFile.
	CreateFileHandlerFunc func(string, Arguments) error

	// WriteContentToHandlerFunc is expected to be provided for as handler for WriteContentTo.
	WriteContentToHandlerFunc func(string, []byte, Arguments) error

	// StreamContentToHandlerFunc is expected to be provided for as handler for StreamContentTo.
	StreamContentToHandlerFunc func(string, io.Reader, Arguments) error

	// CreateDirectoryHandlerFunc is expected to be provided for as handler for CreateDirectory.
	CreateDirectoryHandlerFunc func(string, Arguments) error
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

func createFileDefaultHandler(path string, arg Arguments) (err error) {
	dir, _ := filepath.Split(path)

	di, dErr := os.Stat(dir)
	if dErr != nil {
		if os.IsNotExist(dErr) {
			if !arg.AllowCreationOfDirectoryStructure {
				err = fmt.Errorf("creation of non-existing directory structure is forbidden by current settings: %w", ErrUnresolvableDirectoryStructure)
				return
			}

			if mkdErr := os.MkdirAll(dir, ModeUserReadWriteExecute.asFileMode()); mkdErr != nil {
				err = fmt.Errorf("cannot create directory structure: %w", mkdErr)
				return
			}
		} else {
			err = dErr
			return
		}
	}

	if di != nil && !di.IsDir() {
		err = fmt.Errorf("location structure does not contain valid directory as target: %w", ErrUnresolvableDirectoryStructure)
		return
	}

	fi, fErr := os.Stat(path)
	if fErr != nil && !os.IsNotExist(fErr) {
		err = fErr
		return
	}

	if !arg.AllowOverwrite && fi != nil {
		return ErrFileFound
	}

	if fi != nil && fi.IsDir() {
		err = ErrDirectory
		return
	}

	f, fErr := os.OpenFile(path, os.O_CREATE|os.O_TRUNC, arg.Mode.asFileMode()) //nolint:gosec
	if fErr != nil {
		err = fErr
		return
	}
	defer func() {
		if cErr := f.Close(); cErr != nil {
			err = cErr
		}
	}()

	return nil
}

func writeContentToDefaultHandler(path string, content []byte, arg Arguments) (err error) {
	if oErr := arg.ContentOperation.assetValid(); oErr != nil {
		err = oErr
		return
	}

	flags := os.O_WRONLY | os.O_APPEND
	if arg.ContentOperation.Is(ContentOperationOverwrite) {
		flags = os.O_WRONLY | os.O_TRUNC
	}

	f, fErr := os.OpenFile(path, flags, os.FileMode(0644)) //nolint:gosec
	if fErr != nil {
		if os.IsNotExist(fErr) {
			err = ErrFileNotFound
			return
		}
		err = fErr
		return
	}
	defer func() {
		if cErr := f.Close(); cErr != nil {
			err = cErr
			return
		}
	}()

	n, wErr := f.Write(content)
	if wErr != nil {
		err = wErr
		return
	}

	if n != len(content) {
		err = ErrWriteLengthMismatch
		return
	}

	return
}

func streamContentToDefaultHandler(path string, content io.Reader, arg Arguments) (err error) {
	if oErr := arg.ContentOperation.assetValid(); oErr != nil {
		err = oErr
		return
	}

	flags := os.O_WRONLY | os.O_APPEND
	if arg.ContentOperation.Is(ContentOperationOverwrite) {
		flags = os.O_WRONLY | os.O_TRUNC
	}

	f, fErr := os.OpenFile(path, flags, os.FileMode(0644)) //nolint:gosec
	if fErr != nil {
		if os.IsNotExist(fErr) {
			err = ErrFileNotFound
			return
		}
		err = fErr
		return
	}
	defer func() {
		if cErr := f.Close(); cErr != nil {
			err = cErr
			return
		}
	}()

	if _, wErr := io.Copy(f, content); wErr != nil {
		err = wErr
		return
	}

	return
}

func createDirectoryDefaultHandler(path string, arg Arguments) error {
	dir, _ := filepath.Split(path)

	di, dErr := os.Stat(dir)
	if dErr != nil {
		if os.IsNotExist(dErr) {
			if !arg.AllowCreationOfDirectoryStructure {
				return fmt.Errorf("creation of non-existing directory structure is forbidden by current settings: %w", ErrUnresolvableDirectoryStructure)
			}

			if mkdErr := os.MkdirAll(dir, ModeUserReadWriteExecute.asFileMode()); mkdErr != nil {
				return fmt.Errorf("cannot create directory structure: %w", mkdErr)
			}
		} else {
			return dErr
		}
	}

	if di != nil && !di.IsDir() {
		return fmt.Errorf("location structure does not contain valid directory as target: %w", ErrUnresolvableDirectoryStructure)
	}

	tdi, fErr := os.Stat(path)
	if fErr != nil && !os.IsNotExist(fErr) {
		return fErr
	}

	if tdi != nil {
		if tdi.IsDir() {
			return ErrDirectoryFound
		}
		return ErrFile
	}

	if mkdErr := os.Mkdir(path, arg.Mode.asFileMode()); mkdErr != nil {
		return mkdErr
	}

	return nil
}

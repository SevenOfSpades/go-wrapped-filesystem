package filesystem

import "io"

type defaultFilesystem struct {
	readContentOfHandlerFunc   ReadContentOfHandlerFunc
	streamContentOfHandlerFunc StreamContentOfHandlerFunc
	checkIfExistsHandlerFunc   CheckIfExistsHandlerFunc
}

func newFilesystem(
	readContentOfHandlerFunc ReadContentOfHandlerFunc,
	streamContentOfHandlerFunc StreamContentOfHandlerFunc,
	checkIfExistsHandlerFunc CheckIfExistsHandlerFunc,
) (Filesystem, error) {
	return &defaultFilesystem{
		readContentOfHandlerFunc:   readContentOfHandlerFunc,
		streamContentOfHandlerFunc: streamContentOfHandlerFunc,
		checkIfExistsHandlerFunc:   checkIfExistsHandlerFunc,
	}, nil
}

func (fs *defaultFilesystem) handleReadContentOf(path string) (Content, error) {
	return fs.readContentOfHandlerFunc(path)
}

func (fs *defaultFilesystem) handleStreamContentOf(path string) (io.ReadCloser, error) {
	return fs.streamContentOfHandlerFunc(path)
}

func (fs *defaultFilesystem) handleCheckIfExists(path string) (bool, error) {
	return fs.checkIfExistsHandlerFunc(path)
}

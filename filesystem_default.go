package filesystem

import (
	"io"
)

type defaultFilesystem struct {
	readContentOfHandlerFunc   ReadContentOfHandlerFunc
	streamContentOfHandlerFunc StreamContentOfHandlerFunc
	checkIfExistsHandlerFunc   CheckIfExistsHandlerFunc
	createFileHandlerFunc      CreateFileHandlerFunc
	writeContentToHandlerFunc  WriteContentToHandlerFunc
	streamContentToHandlerFunc StreamContentToHandlerFunc
	createDirectoryHandlerFunc CreateDirectoryHandlerFunc
}

func newFilesystem(
	readContentOfHandlerFunc ReadContentOfHandlerFunc,
	streamContentOfHandlerFunc StreamContentOfHandlerFunc,
	checkIfExistsHandlerFunc CheckIfExistsHandlerFunc,
	createFileHandlerFunc CreateFileHandlerFunc,
	writeContentToHandlerFunc WriteContentToHandlerFunc,
	streamContentToHandlerFunc StreamContentToHandlerFunc,
	createDirectoryHandlerFunc CreateDirectoryHandlerFunc,
) (Filesystem, error) {
	return &defaultFilesystem{
		readContentOfHandlerFunc:   readContentOfHandlerFunc,
		streamContentOfHandlerFunc: streamContentOfHandlerFunc,
		checkIfExistsHandlerFunc:   checkIfExistsHandlerFunc,
		createFileHandlerFunc:      createFileHandlerFunc,
		writeContentToHandlerFunc:  writeContentToHandlerFunc,
		streamContentToHandlerFunc: streamContentToHandlerFunc,
		createDirectoryHandlerFunc: createDirectoryHandlerFunc,
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

func (fs *defaultFilesystem) handleCreateFile(path string, args ...Argument) error {
	arg := &Arguments{
		Mode:                              ModeAllReadWrite,
		DirectoryStructureMode:            ModeAllReadWriteExecute,
		AllowOverwrite:                    false,
		AllowCreationOfDirectoryStructure: false,
	}
	arg.Apply(args)

	return fs.createFileHandlerFunc(path, *arg)
}

func (fs *defaultFilesystem) handleWriteContentTo(path string, content []byte, args ...Argument) error {
	arg := &Arguments{
		ContentOperation: ContentOperationAppend,
	}
	arg.Apply(args)

	return fs.writeContentToHandlerFunc(path, content, *arg)
}

func (fs *defaultFilesystem) handleStreamContentTo(path string, content io.Reader, args ...Argument) error {
	arg := &Arguments{
		ContentOperation: ContentOperationAppend,
	}
	arg.Apply(args)

	return fs.streamContentToHandlerFunc(path, content, *arg)
}

func (fs *defaultFilesystem) handleCreateDirectory(path string, args ...Argument) error {
	arg := &Arguments{
		Mode:                              ModeAllReadWrite,
		DirectoryStructureMode:            ModeAllReadWriteExecute,
		AllowCreationOfDirectoryStructure: false,
	}
	arg.Apply(args)

	return fs.createDirectoryHandlerFunc(path, *arg)
}

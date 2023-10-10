package filesystem

type defaultFilesystem struct {
	readContentOfHandlerFunc   ReadContentOfHandlerFunc
	streamContentOfHandlerFunc StreamContentOfHandlerFunc
}

func newFilesystem(
	readContentOfHandlerFunc ReadContentOfHandlerFunc,
	streamContentOfHandlerFunc StreamContentOfHandlerFunc,
) (Filesystem, error) {
	return &defaultFilesystem{
		readContentOfHandlerFunc:   readContentOfHandlerFunc,
		streamContentOfHandlerFunc: streamContentOfHandlerFunc,
	}, nil
}

func (fs *defaultFilesystem) handleReadContentOf(path string) (Content, error) {
	return fs.readContentOfHandlerFunc(path)
}

func (fs *defaultFilesystem) handleStreamContentOf(path string) (Streamer, error) {
	return fs.streamContentOfHandlerFunc(path)
}

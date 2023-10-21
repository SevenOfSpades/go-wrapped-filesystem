package filesystem

import "github.com/SevenOfSpades/go-just-options"

const (
	optionReadContentOfHandler   options.OptionKey = `read_content_of_handler_func`
	optionStreamContentOfHandler options.OptionKey = `stream_content_of_handler`
	optionCheckIfExistsHandler   options.OptionKey = `check_if_exists_handler`
	optionCreateFileHandler      options.OptionKey = `create_file_handler`
	optionWriteContentToHandler  options.OptionKey = `write_content_to_handler`
	optionStreamContentToHandler options.OptionKey = `stream_content_to_handler`
	optionCreateDirectoryHandler options.OptionKey = `create_directory_handler`
)

// OptionReadContentOfHandler overrides default handler for ReadContentOf.
func OptionReadContentOfHandler(handlerFunc ReadContentOfHandlerFunc) options.Option {
	return func(o options.Options) {
		options.WriteOrPanic[ReadContentOfHandlerFunc](o, optionReadContentOfHandler, handlerFunc)
	}
}

// OptionStreamContentOfHandler overrides default handler for StreamContentOf.
func OptionStreamContentOfHandler(handlerFunc StreamContentOfHandlerFunc) options.Option {
	return func(o options.Options) {
		options.WriteOrPanic[StreamContentOfHandlerFunc](o, optionStreamContentOfHandler, handlerFunc)
	}
}

// OptionCheckIfExistsHandler overrides default handler for CheckIfExists.
func OptionCheckIfExistsHandler(handlerFunc CheckIfExistsHandlerFunc) options.Option {
	return func(o options.Options) {
		options.WriteOrPanic[CheckIfExistsHandlerFunc](o, optionCheckIfExistsHandler, handlerFunc)
	}
}

func OptionCreateFileHandler(handlerFunc CreateFileHandlerFunc) options.Option {
	return func(o options.Options) {
		options.WriteOrPanic[CreateFileHandlerFunc](o, optionCreateFileHandler, handlerFunc)
	}
}

func OptionWriteContentToHandler(handlerFunc WriteContentToHandlerFunc) options.Option {
	return func(o options.Options) {
		options.WriteOrPanic[WriteContentToHandlerFunc](o, optionWriteContentToHandler, handlerFunc)
	}
}

func OptionStreamContentToHandler(handlerFunc StreamContentToHandlerFunc) options.Option {
	return func(o options.Options) {
		options.WriteOrPanic[StreamContentToHandlerFunc](o, optionStreamContentToHandler, handlerFunc)
	}
}

func OptionCreateDirectory(handlerFunc CreateDirectoryHandlerFunc) options.Option {
	return func(o options.Options) {
		options.WriteOrPanic[CreateDirectoryHandlerFunc](o, optionCreateDirectoryHandler, handlerFunc)
	}
}

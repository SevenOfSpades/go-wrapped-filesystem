package filesystem

import "github.com/SevenOfSpades/go-just-options"

const (
	optionReadContentOfHandler   options.OptionKey = `read_content_of_handler_func`
	optionStreamContentOfHandler options.OptionKey = `stream_content_of_handler`
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

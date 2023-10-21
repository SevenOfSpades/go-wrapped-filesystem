package filesystem

type (
	Argument func(args *Arguments)

	Arguments struct {
		AllowCreationOfDirectoryStructure bool
		AllowOverwrite                    bool
		DirectoryStructureMode            Mode
		Mode                              Mode
		ContentOperation                  ContentOperation
	}
)

func (a *Arguments) Apply(args []Argument) {
	for _, x := range args {
		x(a)
	}
}

func WithContentOperation(contentOperation ContentOperation) Argument {
	return func(args *Arguments) {
		args.ContentOperation = contentOperation
	}
}

func WithDirectoryStructureMode(mode Mode) Argument {
	return func(args *Arguments) {
		args.DirectoryStructureMode = mode
	}
}

func WithAllowCreationOfDirectoryStructure(allow bool) Argument {
	return func(args *Arguments) {
		args.AllowCreationOfDirectoryStructure = allow
	}
}

func WithAllowOverwrite(allow bool) Argument {
	return func(args *Arguments) {
		args.AllowOverwrite = allow
	}
}

func WithMode(mode Mode) Argument {
	return func(args *Arguments) {
		args.Mode = mode
	}
}

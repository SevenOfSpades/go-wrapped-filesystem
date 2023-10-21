# go-wrapped-filesystem

Wrapped filesystem makes it easy to implement different handlers / modules for each function.
By default, it will use builtin functions for interacting with filesystem.

## Usage

Use wrapped filesystem with default handlers

```go
package main

import (
	"fmt"
	"log"

	filesystem "github.com/SevenOfSpades/go-wrapped-filesystem"
)

func main() {
	fs, err := filesystem.New()
	if err != nil {
		log.Fatalln(err)
	}

	content, err := filesystem.ReadContentOf(fs, "path/to/file.txt")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print(content.String())
}

```

or implement custom functionality for each function

```go
package main

import (
	"fmt"
	"io"
	"log"

	filesystem "github.com/SevenOfSpades/go-wrapped-filesystem"
)

func main() {
	fs, err := filesystem.New(
		filesystem.OptionReadContentOfHandler(func(path string) (filesystem.Content, error) {
			// do stuff
		}),
		filesystem.OptionStreamContentOfHandler(func(path string) (io.ReadCloser, error) {
			// do stuff
		}),
	)
	if err != nil {
		log.Fatalln(err)
	}

	content, err := filesystem.ReadContentOf(fs, "path/to/file.txt")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print(content.String())
}

```

### List of wrappers

|      Wrapper      | Description                                                                                                                                                                                                                                     |  Works with files  | Works with directories | Package version |      Released      |
|:-----------------:|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|:------------------:|:----------------------:|:---------------:|:------------------:|
|  `ReadContentOf`  | Returns entire content of the file as `filesystem.Content` (`[]byte` with extra functions).                                                                                                                                                     | :white_check_mark: |          :x:           |     v0.0.1      | :white_check_mark: |
| `StreamContentOf` | Returns `io.ReadCloser` attached to the file.<br/>It cannot be guaranteed that all implementations perform actual streaming of the content and won't preload whole file into buffer so please refer to handler's documentation before using it. | :white_check_mark: |          :x:           |     v0.0.1      | :white_check_mark: |
|  `CheckIfExists`  | Returns information (`boolean`) verifying existence of provided path (works for files and directories).                                                                                                                                         | :white_check_mark: |   :white_check_mark:   |     v0.0.2      | :white_check_mark: |
|   `CreateFile`    | Creates file in provided location.                                                                                                                                                                                                              | :white_check_mark: |          :x:           |     v0.0.3      | :white_check_mark: |
| `WriteContentTo`  | Appends/overwrites content to/of provided file.                                                                                                                                                                                                 | :white_check_mark: |          :x:           |     v0.0.3      | :white_check_mark: |
| `StreamContentTo` | Appends/overwrites content to/of provided file from `io.Reader`.                                                                                                                                                                                | :white_check_mark: |          :x:           |     v0.0.3      | :white_check_mark: |
| `CreateDirectory` | Creates new directory at provided location.                                                                                                                                                                                                     |        :x:         |   :white_check_mark:   |     v0.0.3      | :white_check_mark: |

### Arguments

| Argument                                           | Description                                                                                                                               | Wrappers                                                  |             Type              |
|:---------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------|:----------------------------------------------------------|:-----------------------------:|
| `filesystem.WithMode`                              | Changes default mode (`filesystem.ModeAllReadWrite`) for operation in context.                                                            | `filesystem.CreateFile`                                   |       `filesystem.Mode`       |
| `filesystem.WithDirectoryStructureMode`            | Changes default mode (`filesystem.ModeAllReadWriteExecute`) for underlying directory operation in context.                                | `filesystem.CreateFile`                                   |       `filesystem.Mode`       |
| `filesystem.WithAllowCreationOfDirectoryStructure` | Allows for operation in context to create directory structure if not exists.                                                              | `filesystem.CreateFile`                                   |           `boolean`           |
| `filesystem.WithAllowOverwrite`                    | Allows for operation in context to overwrite target if it's the same type.                                                                | `filesystem.CreateFile`                                   |           `boolean`           |
| `filesystem.WithContentOperation`                  | Changes write operation mode between append (`filesystem.ContentOperationAppend`) and overwrite (`filesystem.ContentOperationOverwrite`). | `filesystem.WriteContentTo`, `filesystem.StreamContentTo` | `filesystem.ContentOperation` |

## TODO

- Reading files and directories
    - [x] `ReadContentOf`
    - [x] `StreamContentOf`
    - [X] `CheckIfExists`
    - [ ] `IsFile`
    - [ ] `IsDirectory`
    - [ ] `IsSymlink`
    - [ ] `SizeOf`
    - [ ] `ListFilesIn`
    - [ ] `ReadModeOf`
- Writing to files and directories
    - [x] `CreateFile`
    - [x] `WriteContentTo`
    - [x] `StreamContentTo`
    - [x] `CreateDirectory`
    - [ ] `ChangeModeOf`
- Built-in wrappers
    - [ ] In-Memory *(for tests and stuff)*
    - [ ] HTTP Filesystem *(with server)*
    - [ ] SFTP
    - [ ] S3

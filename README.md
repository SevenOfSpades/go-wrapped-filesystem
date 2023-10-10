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
	"log"

	filesystem "github.com/SevenOfSpades/go-wrapped-filesystem"
)

func main() {
	fs, err := filesystem.New(
		filesystem.OptionReadContentOfHandler(func(path string) (filesystem.Content, error) {
			// do stuff
		}),
		filesystem.OptionStreamContentOfHandler(func(path string) (filesystem.Streamer, error) {
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
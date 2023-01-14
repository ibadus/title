# Title

> A small package for changing the terminal title in Golang.
> Supporting Windows, Linux and Mac OS.

## Installation

```bash
go get github.com/ibadus/title
```

## Usage

```go
package main

import (
	"github.com/ibadus/title"
)

func main() {
	title.Set("Hello World!")
}
```

You can also use the `Setf` function to include formatted strings:

```go
package main

import (
	"github.com/ibadus/title"
)

func main() {
	title.Setf("Hello %s!", "World)
}
```

# License
This package is released under the MIT license. See LICENSE for more information.
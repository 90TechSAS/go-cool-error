# go-cool-error
An tiny and beautiful error printer

## Installation

First, get ryanuber/columnize:
```bash
go get github.com/ryanuber/columnize
```

And then, get go-cool-error:
```bash
go get github.com/90TechSAS/go-cool-error
```

## Example
```go 
package main

import (
	"errors"
	"fmt"
	"github.com/90TechSAS/go-cool-error"
)

func main() {
	err := A()
	if err.Err != nil {
		fmt.Println(err.GetStack())
	} else {
		fmt.Println("All Good")
	}
}

func A() (e cerror.Error) {
	e = B()
	return e
}

func B() (e cerror.Error) {
	e = C()
	return e
}

func C() (e cerror.Error) {
	e.Return(errors.New("Error message #1"), []string{"Error", "message"}, 2, true)
	return e
}
```

Output:

```
Error: Error message #1 ([Error message] 2 true)
  main.C     /tmp/cerror/examples/main.go:30
  main.B     /tmp/cerror/examples/main.go:25
  main.A     /tmp/cerror/examples/main.go:20
  main.main  /tmp/cerror/examples/main.go:1
```
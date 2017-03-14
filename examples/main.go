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

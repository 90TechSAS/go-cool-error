package cerror

import (
	"errors"
	"fmt"
	"runtime"
	"strings"

	"github.com/ryanuber/columnize"
)

/*
	Error struct

	Err:	the error
	stack:	the formatted stack
*/
type Error struct {
	Err   error
	stack string
}

/*
	Return the stack
*/
func (e *Error) GetStack() string {
	return e.stack
}

/*
	Used for returning an error
*/
func (e *Error) Return(err error, msg ...interface{}) {
	var sprint string
	if err == nil {
		return
	}
	e.Err = err
	if len(msg) > 0 {
		sprint += " ("
		for _, token := range msg {
			sprint += fmt.Sprint(token, " ")
		}
		sprint = sprint[:len(sprint)-1] // Don't show the last space of sprint
		sprint += ")"
	}
	e.stack = "Error: " + err.Error() + sprint + "\n" + getStack() // Concat the error message and the stack
	e.stack = e.stack[:len(e.stack)-1]                             // Don't show the last end of line of the stack
}

/*
	Return the stack
*/
func getStack() string {
	var stack []string
	for i := 3; ; i++ { // Begin at 3 to not print the loger stack itself
		line, err := getStackLine(i)
		if err != nil { // Break when the
			break
		}
		stack = append(stack, line)
	}
	return columnize.SimpleFormat(stack)
}

/*
	Return the stack given line
*/
func getStackLine(i int) (string, error) {
	pc, file, line, ok := runtime.Caller(i)
	// Get the function only, not the full path (in form of package.function)
	FullFunctionName := runtime.FuncForPC(pc).Name()
	functionNameArr := strings.Split(FullFunctionName, "/")
	functionName := functionNameArr[len(functionNameArr)-1]
	if ok == false || functionName == "runtime.main" {
		return "", errors.New("Stack terminated")
	}
	return fmt.Sprintf("|%s|%s:%d", functionName, file, line), nil
}

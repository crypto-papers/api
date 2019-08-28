package errors

import (
	"errors"
	"fmt"
	"log"
	"runtime"
)

var (
	InternalServerError = generateError("Internal server error.")
)

func generateError(err string) error {
	return errors.New(err)
}

func DebugError(err_ error, args ...interface{}) string {
	programCounter, file, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(programCounter)
	msg := fmt.Sprintf("[%s: %s %d] %s, %s", file, fn.Name(), line, err_, args)
	log.Println(msg)
	return msg
}

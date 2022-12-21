package utils

import (
	"fmt"
	"time"
)

type LoggingUtils struct{}

func (LoggingUtils) Debug(v ...any) {
	t := time.Now().Local().Format("02/01/2006 15:04:05")
	fmt.Printf("[\x1b[34;1m%s\x1b[0m] [\x1b[35;1mDEBUG\x1b[0m] ", t)
	fmt.Println(v...)
}

func (LoggingUtils) Info(v ...any) {
	t := time.Now().Local().Format("02/01/2006 15:04:05")
	fmt.Printf("[\x1b[34;1m%s\x1b[0m] [\x1b[32;1mINFO\x1b[0m] ", t)
	fmt.Println(v...)
}

func (LoggingUtils) Warn(v ...any) {
	t := time.Now().Local().Format("02/01/2006 15:04:05")
	fmt.Printf("[\x1b[34;1m%s\x1b[0m] [\x1b[33;1mWARNING\x1b[0m] ", t)
	fmt.Println(v...)
}

func (LoggingUtils) Error(v ...any) {
	t := time.Now().Local().Format("02/01/2006 15:04:05")
	fmt.Printf("[\x1b[34;1m%s\x1b[0m] [\x1b[31;1mERROR\x1b[0m] ", t)
	fmt.Println(v...)
}

var Logging = &LoggingUtils{}

package logger

import (
	"fmt"
)

const (
	DebugLevel   = 0
	InfoLevel    = 1
	WarnLevel    = 2
	ErrorLevel   = 3
	FatalLevel   = 4
	UnknownLevel = 5
)

var CurrentLevel int

func Debug(message string)   { Out(DebugLevel, message) }
func Info(message string)    { Out(InfoLevel, message) }
func Warn(message string)    { Out(WarnLevel, message) }
func Error(message string)   { Out(ErrorLevel, message) }
func Fatal(message string)   { Out(FatalLevel, message) }
func Unknown(message string) { Out(UnknownLevel, message) }

func Out(level int, message string) {
	if level >= CurrentLevel {
		fmt.Println(message)
	}
}

func init() {
	CurrentLevel = InfoLevel
}

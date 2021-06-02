package core

import (
	"fmt"
	"os"
)

var minLevel uint8

var levels map[string]uint8 = map[string]uint8{
	"ALL":     0,
	"DEBUG":   1,
	"LOG":     2,
	"INFO":    3,
	"WARN":    4,
	"ERROR":   5,
	"SILENCE": 6,
}

func init() {
	logLevel := os.Getenv("LOGGING_LEVEL")
	if logLevel == "" {
		minLevel = 1
		fmt.Println("[INFO]  Setting default log level to INFO")
		return
	}

	level, levelOk := levels[logLevel]
	if !levelOk {
		fmt.Println("[INFO]  Setting default log level INFO")
		return
	}

	minLevel = level
}

// Warn prints a [WARN] msg
func Warn(msg string, args ...interface{}) {
	if minLevel > levels["WARN"] {
		return
	}

	full := "[WARN]  " + msg + "\n"
	fmt.Printf(full, args...)
}

// Debug prints a [DEBUG] msg
func Debug(msg string, args ...interface{}) {
	if minLevel > levels["DEBUG"] {
		return
	}

	full := "[DEBUG] " + msg + "\n"
	fmt.Printf(full, args...)
}

// Log prints an [INFO] Log
func Log(msg string, args ...interface{}) {
	if minLevel > levels["INFO"] {
		return
	}

	full := "[INFO]  " + msg + "\n"
	fmt.Printf(full, args...)
}

// Error prints a [ERROR] msg
func Error(msg string, args ...interface{}) {
	if minLevel > levels["ERROR"] {
		return
	}

	full := "[ERROR] " + msg + "\n"
	fmt.Printf(full, args...)
}

package util

import (
	"log"
	"os"
)

// Coloured labels
// https://www.dolthub.com/blog/2024-02-23-colors-in-golang/
const (
	reset = "\033[0m"
	error = "\033[31m" + "ERROR: " + reset
	warn  = "\033[33m" + "WARN: " + reset
	info  = "\033[34m" + "INFO: " + reset
	debug = "\033[36m" + "DEBUG: " + reset
)

// Custom API logger
type ILogger interface {
	Error(v ...interface{})
	Warn(v ...interface{})
	Info(v ...interface{})
	Debug(v ...interface{})
}

// Logging level as integer
type Level int

type Logger struct {
	level    Level
	levelMap map[string]Level
	debug    *log.Logger
	info     *log.Logger
	warn     *log.Logger
	error    *log.Logger
}

// Create new logger
//
// Log level determines which logs will show
// ie: Setting DEBUG will render all log levels
func NewLogger(logger *log.Logger) *Logger {
	return &Logger{
		level:    3,
		levelMap: map[string]Level{"ERROR": 0, "WARN": 1, "INFO": 2, "DEBUG": 3},
		error:    log.New(os.Stdout, error, log.LstdFlags),
		warn:     log.New(os.Stdout, warn, log.LstdFlags),
		info:     log.New(os.Stdout, info, log.LstdFlags),
		debug:    log.New(os.Stdout, debug, log.LstdFlags),
	}
}

// Set the current log level DEBUG | INFO | WARN | ERROR
func (l *Logger) SetLogLevel(level string) {
	newLevel, ok := l.levelMap[level]

	if ok {
		l.Info("Setting log level to:", level)
		l.level = newLevel
		return
	}

	l.Warn("Log level not supported:", level)
}

// Debug logs for development
func (l *Logger) Debug(v ...interface{}) {
	if l.level >= l.levelMap["DEBUG"] {
		l.debug.Println(v...)
	}
}

// Info logs for additional information
func (l *Logger) Info(v ...interface{}) {
	if l.level >= l.levelMap["INFO"] {
		l.info.Println(v...)
	}
}

// Warn logs for potential errors
func (l *Logger) Warn(v ...interface{}) {
	if l.level >= l.levelMap["WARN"] {
		l.warn.Println(v...)
	}
}

// Error logs
func (l *Logger) Error(v ...interface{}) {
	if l.level >= l.levelMap["ERROR"] {
		l.error.Println(v...)
	}
}

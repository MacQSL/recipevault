package server

import (
	"log"
	"os"
)

// Logging level as integer
type Level int

type ILogger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
}

type Logger struct {
	level    Level
	levelMap map[string]Level
	debug    *log.Logger
	info     *log.Logger
	warn     *log.Logger
	error    *log.Logger
}

// Create new Logger
//
// Log level determines which logs will show
// ie: Setting DEBUG will render all log levels
func NewLogger(logger *log.Logger) *Logger {
	return &Logger{
		level:    0,
		levelMap: map[string]Level{"ERROR": 0, "WARN": 1, "INFO": 2, "DEBUG": 3},
		error:    log.New(os.Stdout, "ERROR: ", log.LstdFlags),
		warn:     log.New(os.Stdout, "WARN: ", log.LstdFlags),
		info:     log.New(os.Stdout, "INFO: ", log.LstdFlags),
		debug:    log.New(os.Stdout, "DEBUG: ", log.LstdFlags),
	}
}

// Set log level - renders all logs below level aswell
// Excluding error logs which always render
func (l *Logger) setLogLevel(level string) {
	newLevel, ok := l.levelMap[level]

	if ok {
		l.info.Printf("Setting log level to: %s\n", level)
		l.level = newLevel
		return
	}

	l.warn.Printf("Log level not supported: %s\n", level)
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
		l.info.Println(v...)
	}
}

// Error logs - always
func (l *Logger) Error(v ...interface{}) {
	if l.level >= l.levelMap["ERROR"] {
		l.error.Println(v...)
	}
}

package util

import (
	"log"
	"os"
)

// Coloured labels
// https://www.dolthub.com/blog/2024-02-23-colors-in-golang/
const (
	resetTag = "\033[0m"
	fatalTag = "\033[35m" + "FATAL: " + resetTag
	errorTag = "\033[31m" + "ERROR: " + resetTag
	warnTag  = "\033[33m" + "WARN: " + resetTag
	infoTag  = "\033[34m" + "INFO: " + resetTag
	debugTag = "\033[36m" + "DEBUG: " + resetTag
)

// Custom API logger
type ILogger interface {
	Fatal(v ...interface{})
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
	fatal    *log.Logger
	error    *log.Logger
	warn     *log.Logger
	info     *log.Logger
	debug    *log.Logger
}

// Create new logger
//
// Log level determines which logs will show
// ie: Setting DEBUG will render all log levels
func NewLogger() *Logger {
	return &Logger{
		level:    3,
		levelMap: map[string]Level{"FATAL": 0, "ERROR": 1, "WARN": 2, "INFO": 3, "DEBUG": 4},
		fatal:    log.New(os.Stdout, fatalTag, log.LstdFlags),
		error:    log.New(os.Stdout, errorTag, log.LstdFlags),
		warn:     log.New(os.Stdout, warnTag, log.LstdFlags),
		info:     log.New(os.Stdout, infoTag, log.LstdFlags),
		debug:    log.New(os.Stdout, debugTag, log.LstdFlags),
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

// Loggers in order of display precedence

// Fatal logs - calls os.exit(1) to kill program
func (l *Logger) Fatal(v ...interface{}) {
	if l.level >= l.levelMap["FATAL"] {
		l.fatal.Println(v...)
		os.Exit(1)
	}
}

// Error logs
func (l *Logger) Error(v ...interface{}) {
	if l.level >= l.levelMap["ERROR"] {
		l.error.Println(v...)
	}
}

// Warn logs for potential errors
func (l *Logger) Warn(v ...interface{}) {
	if l.level >= l.levelMap["WARN"] {
		l.warn.Println(v...)
	}
}

// Info logs for additional information
func (l *Logger) Info(v ...interface{}) {
	if l.level >= l.levelMap["INFO"] {
		l.info.Println(v...)
	}
}

// Debug logs for development
func (l *Logger) Debug(v ...interface{}) {
	if l.level >= l.levelMap["DEBUG"] {
		l.debug.Println(v...)
	}
}

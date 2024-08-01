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
func NewLogger(logger *log.Logger) *Logger {
	return &Logger{
		levelMap: map[string]Level{"DEBUG": 0, "INFO": 1, "WARN": 2, "ERROR": 3},
		debug:    log.New(os.Stdout, "DEBUG: ", log.LstdFlags),
		info:     log.New(os.Stdout, "INFO: ", log.LstdFlags),
		warn:     log.New(os.Stdout, "WARN: ", log.LstdFlags),
		error:    log.New(os.Stdout, "ERROR: ", log.LstdFlags),
	}
}

// Set log level - renders all logs below level aswell
func (l *Logger) SetLogLevel(level string) {
	newLevel, ok := l.levelMap[level]

	if ok {
		l.debug.Println("Setting log level to: %s", level)
		l.level = newLevel
		return
	}

	l.debug.Println("Log level not supported: ", level)
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

// Error logs
func (l *Logger) Error(v ...interface{}) {
	if l.level == l.levelMap["ERROR"] {
		l.error.Println(v...)
	}
}

package util

import (
    "log"
    "os"
)

type Logger struct {
    *log.Logger
}

var logger *Logger

// Initialize the logger
func init() {
    logLevel := os.Getenv("LOG_LEVEL")
    if logLevel == "" {
        logLevel = "INFO" // Default log level
    }

    logger = NewLogger(logLevel)
}

// NewLogger creates a new logger instance with the specified log level
func NewLogger(level string) *Logger {
    flags := log.Ldate | log.Ltime | log.Lshortfile
    logger := log.New(os.Stdout, "", flags)

    return &Logger{logger}
}

// GetLogger returns the logger instance
func GetLogger() *Logger {
    return logger
}

// InfoGeneral logs informational general messages
func (l *Logger) InfoGeneral(v ...interface{}) {
    l.SetPrefix("☑ INFO: ")
    l.Println(v...)
}
// InfoSuccess logs informational success messages
func (l *Logger) InfoSuccess(v ...interface{}) {
    l.SetPrefix("☑ SUCCESS: ")
    l.Println(v...)
}

// Warning logs warning messages
func (l *Logger) Warning(v ...interface{}) {
    l.SetPrefix("⚠ WARNING: ")
    l.Println(v...)
}

// Error logs error messages
func (l *Logger) Error(v ...interface{}) {
    l.SetPrefix("🔴ERROR: ")
    l.Println(v...)
}

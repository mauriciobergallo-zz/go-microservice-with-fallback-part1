package logging

import (
	"encoding/json"
	"fmt"
	"time"
)

// Service unified
type Service interface {
	Debug(...string)
	Info(...string)
	Warn(...string)
	Error(...string)
}

type defaultLogging struct {
	logLevel string
}

// NewStdoutLogging constructor of the logging service
func NewStdoutLogging(ll string) Service {
	return &defaultLogging{logLevel: ll}
}

// Logs a DEBUG level message, the parameters are ClassName, Message...
func (dl *defaultLogging) Debug(p ...string) {
	if dl.logLevel == "DEBUG" {
		dl.print("DEBUG", p[0], p[1], "", 0)
	}
}

// Logs an INFO level message, the parameters are ClassName, Message...
func (dl *defaultLogging) Info(p ...string) {
	if dl.logLevel == "DEBUG" || dl.logLevel == "INFO" {
		dl.print("INFO", p[0], p[1], "", 0)
	}
}

// Logs a WARN level message, the parameters are ClassName, Message...
func (dl *defaultLogging) Warn(p ...string) {
	if dl.logLevel == "DEBUG" || dl.logLevel == "INFO" || dl.logLevel == "WARN" {
		dl.print("WARN", p[0], p[1], "", 0)
	}
}

// Logs an ERROR level message, the parameters are ClassName, Message, Exception...
func (dl *defaultLogging) Error(p ...string) {
	if dl.logLevel == "DEBUG" || dl.logLevel == "INFO" || dl.logLevel == "WARN" || dl.logLevel == "ERROR" {
		dl.print("ERROR", p[0], p[1], p[2], 0)
	}
}

func (dl *defaultLogging) print(logLevel string, class string, msg string, exc string, et int64) {
	l := Log{
		CorrelationID: "",
		ClassName:     class,
		TimeStamp:    time.Now(),
		Message:       msg,
		Exception:     exc,
		LogLevel:      logLevel,
	}

	pp, err := json.Marshal(l)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(pp))
}
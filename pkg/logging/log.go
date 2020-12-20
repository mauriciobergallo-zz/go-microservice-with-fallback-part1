package logging

import "time"

type Log struct {
	CorrelationID string    `json:"correlationId"`
	ClassName     string    `json:"className"`
	TimeStamp    time.Time `json:"timestamp"`
	Message       string    `json:"message"`
	Exception     string    `json:"exception"`
	LogLevel      string    `json:"logLevel"`
	ExecutionTime int64     `json:"executionTime"`
}
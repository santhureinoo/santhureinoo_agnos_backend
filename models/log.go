package models

import "time"

type LogEntry struct {
	Method     string
	Path       string
	StatusCode int
	Duration   time.Duration
}

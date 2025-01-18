package Mongomodels

import "time"

type UserLogs struct {
	Goroutine int       `json:"goroutine"`
	Timestamp time.Time `json:"time"`
	Message   string    `json:"message"`
}

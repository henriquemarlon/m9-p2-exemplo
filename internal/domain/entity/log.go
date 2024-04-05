package entity

import (
	"time"
)

type LogRepository interface {
	CreateLog(data *Log) error
}

type Log struct {
	ID        string    `json:"sensor_id"`
	Unit      string    `json:"unit"`
	Level     float64   `json:"level"`
	Timestamp time.Time `json:"timestamp"`
}

func NewLog(id string, unit string, level float64, timestamp time.Time) *Log {
	return &Log{
		ID:        id,
		Unit:      unit,
		Level:     level,
		Timestamp: timestamp,
	}
}
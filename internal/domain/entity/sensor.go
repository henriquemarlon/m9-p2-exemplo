package entity

import (
	"math"
	"math/rand"
	"time"
)

type SensorRepository interface {
	FindAllSensors() ([]*Sensor, error)
}

type Sensor struct {
	ID     string                 `json:"_id"`
	Unit   string                 `json:"unit"`
	Params map[string]interface{} `json:"params"`
}

type SensorPayload struct {
	ID        string    `json:"sensor_id"`
	Unit      string    `json:"unit"`
	Level     float64   `json:"level"`
	Timestamp time.Time `json:"timestamp"`
}

func Entropy(min float64, max float64) float64 {
	rand.NewSource(time.Now().UnixNano())
	return math.Round(float64(rand.Float64()*(max-min) + min))
}

func NewSensorPayload(id string, unit string, params map[string]interface{}, timestamp time.Time) (*SensorPayload, error) {
	min, ok := params["min"].(float64)
	if !ok {
		panic("min value not found or not a float64")
	}
	max, ok := params["max"].(float64)
	if !ok {
		panic("max value not found or not a float64")
	}

	value := Entropy(min, max)
	percent := (value - min) / (max - min) * 100

	return &SensorPayload{
		ID:        id,
		Unit:      unit,
		Level:     percent,
		Timestamp: timestamp,
	}, nil
}

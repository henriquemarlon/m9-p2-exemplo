package usecase

import (
	"log"
	"time"

	"github.com/henriquemarlon/m9-p2/internal/domain/entity"
)

type CreateLogUseCase struct {
	LogRepository entity.LogRepository
}

type CreateLogInputDTO struct {
	Sensor_ID        string    `json:"sensor_id"`
	Unit      string    `json:"unit"`
	Level     float64   `json:"level"`
	Timestamp time.Time `json:"timestamp"`
}

func NewLogUseCase(logRepository entity.LogRepository) *CreateLogUseCase {
	return &CreateLogUseCase{LogRepository: logRepository}
}

func (c *CreateLogUseCase) Execute(input CreateLogInputDTO) error {
	logData := entity.NewLog(input.Sensor_ID, input.Unit, input.Level, input.Timestamp)
	err := c.LogRepository.CreateLog(logData)
	if err != nil {
		log.Printf("Error creating sensor log: %v", err)
	}
	return nil
}

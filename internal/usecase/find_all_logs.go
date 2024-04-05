package usecase

import (
	"github.com/henriquemarlon/m9-p2/internal/domain/entity"
	"time"
)

type FindAllLogsUseCase struct {
	LogRepository entity.LogRepository
}

type FindAllLogsInputDTO struct {
	ID        string    `json:"sensor_id"`
	Unit      string    `json:"unit"`
	Level     float64   `json:"level"`
	Timestamp time.Time `json:"timestamp"`
}


func NewFindAllLogsUseCase(logRepository entity.LogRepository) *FindAllLogsUseCase {
	return &FindAllLogsUseCase{LogRepository: logRepository}
}

func (f *FindAllLogsUseCase) Execute() ([]FindAllLogsInputDTO, error) {
	logs, err := f.LogRepository.FindAllLogs()
	if err != nil {
		return nil, err
	}
	var output []FindAllLogsInputDTO
	for _, log := range logs {
		output = append(output, FindAllLogsInputDTO{
			ID:        log.ID,
			Unit:      log.Unit,
			Level:     log.Level,
			Timestamp: log.Timestamp,
		})
	}
	return output, nil
}
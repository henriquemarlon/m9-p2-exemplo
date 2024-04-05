package usecase

import (
	"github.com/henriquemarlon/m9-p2/internal/domain/entity"
)

type FindAllSensorsUseCase struct {
	SensorRepository entity.SensorRepository
}

type FindAllSensorsOutputDTO struct {
	ID        string                 `json:"sensor_id"`
	Unit      string                 `json:"unit"`
	Params    map[string]interface{} `json:"params"`
}

func NewFindAllSensorsUseCase(sensorRepository entity.SensorRepository) *FindAllSensorsUseCase {
	return &FindAllSensorsUseCase{SensorRepository: sensorRepository}
}

func (f *FindAllSensorsUseCase) Execute() ([]FindAllSensorsOutputDTO, error) {
	sensors, err := f.SensorRepository.FindAllSensors()
	if err != nil {
		return nil, err
	}
	var output []FindAllSensorsOutputDTO
	for _, sensor := range sensors {
		output = append(output, FindAllSensorsOutputDTO{
			ID:        sensor.ID,
			Unit:      sensor.Unit,
			Params:    sensor.Params,
		})
	}
	return output, nil
}

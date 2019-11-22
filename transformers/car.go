package transformers

import (
	"hello/dtos"
	"hello/models"
)

func TransformToCarsDTO(car models.Car) (carDTO dtos.CarDTO) {
	carDTO = dtos.CarDTO{
		ID:             car.ID,
		Owner:          car.Owner,
		Color:          car.Color,
		NumberOfWheels: car.NumberOfWheels,
	}
	return
}

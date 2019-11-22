package services

import (
	"fmt"
	"hello/daos"
	"hello/dtos"
	"hello/transformers"
)

type CarService interface {
	//TransformToCarDTO(car models.Car) (carDTO dtos.CarDTO)
	GetAll() (carsDTO []dtos.CarDTO, err error)
	GetCarById(id int) (carDTO dtos.CarDTO, err error)
}

type CarServiceImpl struct {
	CarDAO daos.CarDAOImpl
}

func (carService CarServiceImpl) GetAll() (carsDTO []dtos.CarDTO, err error) {
	cars, err := carService.CarDAO.GetAll()
	if err != nil {
		fmt.Println("cannot get cars in service" + err.Error())
	}

	carsDTO = make([]dtos.CarDTO, 0)
	for _, car := range cars {
		carsDTO = append(carsDTO, transformers.TransformToCarsDTO(car))
	}
	return
}

func (carService CarServiceImpl) GetCarById(id int) (carDTO dtos.CarDTO, err error) {
	car, err := carService.CarDAO.GetCarById(id)
	if err != nil {
		fmt.Println("cannot get car in service" + err.Error())
		return
	}
	carDTO = transformers.TransformToCarsDTO(car)
	return
}

func NewCarService(carDAOImpl daos.CarDAOImpl) (carService CarServiceImpl) {
	carService.CarDAO = carDAOImpl
	return carService
}
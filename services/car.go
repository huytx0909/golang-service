package services

import (
	"fmt"

	"hello/daos"
	"hello/dtos"
	"hello/models"
	"hello/transformers"
)

type CarService interface {
	GetAll() (carsDTO []dtos.CarDTO, err error)
	GetCarById(id int) (carDTO dtos.CarDTO, err error)
	CreateCar(carDTO dtos.CarDTO) (err error)
	DeleteCarById(id int) (err error)
	UpdateCarById(id uint, carDTO dtos.CarDTO) (err error)
}

type CarServiceImpl struct {
	CarDAO daos.CarDAOImpl
}

func NewCarService(carDAOImpl daos.CarDAOImpl) (carService CarServiceImpl) {
	carService.CarDAO = carDAOImpl
	return carService
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

func (carService CarServiceImpl) CreateCar(carDTO dtos.CarDTO) (err error) {
	car := models.Car{
		Owner:          carDTO.Owner,
		Color:          carDTO.Color,
		NumberOfWheels: carDTO.NumberOfWheels,
	}

	err = carService.CarDAO.CreateCar(car)
	if err != nil {
		fmt.Println("cannot create car" + err.Error())
	}
	return
}

func (carService CarServiceImpl) DeleteCarById(id int) (err error) {
	err = carService.CarDAO.DeleteCarById(id)
	if err != nil {
		fmt.Println("cannot delete in service")
	}
	return
}

func (carService CarServiceImpl) UpdateCarById(id uint, carDTO dtos.CarDTO) (err error) {
	carModel := models.Car{
		Owner:          carDTO.Owner,
		Color:          carDTO.Color,
		NumberOfWheels: carDTO.NumberOfWheels,
	}
	err = carService.CarDAO.UpdateCarById(id, carModel)
	if err != nil {
		fmt.Println("cannot update in service")
	}
	return
}

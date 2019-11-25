package daos

import (
	"github.com/jinzhu/gorm"

	"hello/models"
)

// CarDAOs
type CarDAO interface {
	GetAll() (cars []models.Car, err error)
	GetCarById(id int) (car models.Car, err error)
	CreateCar(car models.Car) (err error)
	UpdateCarById(id int, newCar models.Car) (err error)
	DeleteCarById(id int) (err error)
}

// DAOImpl
type CarDAOImpl struct {
	db *gorm.DB
}

// NewCarDAO
func NewCarDAO(db *gorm.DB) (carDAO CarDAOImpl) {
	carDAO.db = db
	return
}

func (carDAO *CarDAOImpl) GetAll() (cars []models.Car, err error) {
	cars = make([]models.Car, 0)
	carDAO.db.Find(&cars)
	return
}

func (carDAO *CarDAOImpl) GetCarById(id int) (car models.Car, err error) {
	car = models.Car{}
	carDAO.db.First(&car, id)
	return
}

func (carDAO *CarDAOImpl) CreateCar(car models.Car) (err error) {
	carDAO.db.Create(&car)
	return
}

func (carDAO *CarDAOImpl) DeleteCarById(id int) (err error) {
	car := models.Car{}
	carDAO.db.First(&car, id)
	if car.ID == 0 {
		return
	}
	carDAO.db.Delete(&car)
	return
}

func (carDAO *CarDAOImpl) UpdateCarById(id uint, newCar models.Car) (err error) {
	car := models.Car{}
	carDAO.db.First(&car, id)
	if car.ID == 0 {
		return
	}
	carDAO.db.Model(&car).Updates(models.Car{
		Owner:          newCar.Owner,
		Color:          newCar.Color,
		NumberOfWheels: newCar.NumberOfWheels,
	})

	return
}

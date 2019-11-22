package daos

import (
	"github.com/jinzhu/gorm"
	"hello/models"
)

// CarDAOs
type CarDAO interface {
	GetAll() (cars []models.Car, err error)
	GetCarById(id int) (car models.Car, err error)
	//CreateCar(carDTO dtos.CarDTO) (err error)
	//UpdateCar(id int, carDTO dtos.CarDTO) (err error)
	//DeleteCarById(id int) (err error)
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
	carDAO.db.First(&car, id)
	return
}

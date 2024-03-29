package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"hello/dtos"
	"hello/services"
)

type CarController struct {
	CarService services.CarService
}

func NewCarController(carService services.CarService) (carController CarController) {
	carController.CarService = carService
	return
}

// Get all cars
func (carController *CarController) GetAll(c *gin.Context) {
	cars, err := carController.CarService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   cars,
	})
}

// Get car by id
func (carController *CarController) GetCarById(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	car, err := carController.CarService.GetCarById(idInt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   car,
	})
}

// Create car
func (carController *CarController) CreateCar(c *gin.Context) {
	carDTO := new(dtos.CarDTO)
	err := c.BindJSON(carDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	err = carController.CarService.CreateCar(*carDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": http.StatusCreated,
	})
}

// Delete car by id
func (carController *CarController) DeleteCarById(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	err = carController.CarService.DeleteCarById(idInt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
	})
}

// Update car by id
func (carController *CarController) UpdateCarById(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	carDTO := dtos.CarDTO{}
	err = c.BindJSON(&carDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	err = carController.CarService.UpdateCarById(idInt, carDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"status": http.StatusAccepted,
	})
}

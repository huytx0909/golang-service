package controllers

import (
	"github.com/gin-gonic/gin"
	"hello/services"
	"net/http"
	"strconv"
)

type CarController struct {
	CarServiceImpl services.CarServiceImpl
}

func NewCarController(carService services.CarServiceImpl) (carController CarController) {
	carController.CarServiceImpl = carService
	return
}

func (carController *CarController) GetAll(c *gin.Context) {
	cars, err := carController.CarServiceImpl.GetAll()
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

	car, err := carController.CarServiceImpl.GetCarById(idInt)

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


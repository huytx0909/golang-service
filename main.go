package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/lib/pq"

	"hello/controllers"
	"hello/daos"
	"hello/services"
)

var db *gorm.DB

func init() {
	// Open DB connection
	var err error
	db, err = gorm.Open("mysql", "huytx:huytx@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("error connecting to db" + err.Error())
	}

}

func main() {
	defer db.Close()
	// Init application
	NewCarApplication()

	// Migrations
	driver, err := mysql.WithInstance(db.DB(), &mysql.Config{})

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "mysql", driver)
	if err != nil {
		panic("error migrating " + err.Error())
	}

	err = m.Up()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer m.Close()

}

func NewCarApplication() {

	// Init DAOs
	carDAO := daos.NewCarDAO(db)

	// Init Services
	carService := services.NewCarService(&carDAO)

	// Init Controllers
	controller := controllers.NewCarController(&carService)

	router := gin.Default()

	v1 := router.Group("/api/v1/car")
	{
		v1.POST("/", controller.CreateCar)
		v1.GET("/", controller.GetAll)
		v1.GET("/:id", controller.GetCarById)
		v1.PUT("/:id", controller.UpdateCarById)
		v1.DELETE("/:id", controller.DeleteCarById)
	}
	router.Run()
}

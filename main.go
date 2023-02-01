package main

import (
	"StudentAPI/config"
	"StudentAPI/controller"
	"StudentAPI/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// Create HTTP server
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"hello": "world",
		})
	})

	// Connect To Database
	config.DatabaseInit()
	gorm := config.DB()
	gorm.AutoMigrate(&model.Student{}, &model.Class{}, &model.Mark{})

	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}
	defer dbGorm.Close()

	//Student API routes
	studentRoute := e.Group("/student")
	studentRoute.POST("/", controller.CreateStudent)
	studentRoute.GET("/:id", controller.GetStudent)
	studentRoute.PUT("/:id", controller.UpdateStudent)
	studentRoute.DELETE("/:id", controller.DeleteStudent)
	studentRoute.GET("/best_student/class/:id", controller.GetBestStudentInClass)

	//Student API routes
	classRoute := e.Group("/class")
	classRoute.POST("/", controller.CreateClass)
	classRoute.GET("/:id", controller.GetClass)
	classRoute.PUT("/:id", controller.UpdateClass)
	classRoute.DELETE("/:id", controller.DeleteClass)
	classRoute.GET("/:id/students", controller.GetClassStudents)

	//Route to fetch student who has best note in x class

	e.Logger.Fatal(e.Start(":8080"))
}

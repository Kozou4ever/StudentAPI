package main

import (
	"StudentAPI/config"
	"StudentAPI/controller"
	"StudentAPI/model"

	"github.com/labstack/echo/v4"
)

func main() {
	// Create HTTP server
	e := echo.New()

	// Connect To Database
	config.DatabaseInit()
	gorm := config.DB()

	// Migrate the schema
	gorm.AutoMigrate(&model.Student{}, &model.Class{}, &model.Mark{})

	// Close database connection
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

	//Start server in 8080 port
	e.Logger.Fatal(e.Start(":8080"))
}

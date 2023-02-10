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
	studentRoute.GET("/:id", controller.GetStudentDetails)
	studentRoute.PUT("/:id", controller.UpdateStudentDetails)
	studentRoute.DELETE("/:id", controller.DeleteStudent)
	studentRoute.GET("/rank-student/:rank", controller.RankStudents)
	studentRoute.GET("/best-student", controller.BestStudents)

	//Student API routes
	classRoute := e.Group("/class")
	classRoute.POST("/", controller.CreateClass)
	classRoute.GET("/:id", controller.GetClassDetails)
	classRoute.PUT("/:id", controller.UpdateClassDetails)
	classRoute.DELETE("/:id", controller.DeleteClass)
	classRoute.GET("/:id/top-student", controller.GetTopStudentsClass)
	//Start server in 8080 port
	e.Logger.Fatal(e.Start(":8080"))
}

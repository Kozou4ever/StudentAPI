package main

import (
	"StudentAPI/database"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func main() {
	e := echo.New()
	database.Connect()
	postgreDB, _ := database.DB.DB()
	defer postgreDB.Close()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/student", saveStudent)
	//e.GET("/student/:id", getStudent)
	//e.PUT("/student/:id", updateStudent)
	//e.DELETE("/student/:id", deleteStudent)

	e.POST("/class", saveClass)

	e.GET("/class/:id/students", searchStudentsByClasss)

	e.Logger.Fatal(e.Start(":1323"))
}

func saveStudent(c echo.Context) error {
	student := database.Student{}
	if err := c.Bind(&student); err != nil {
		return err
	}
	database.DB.Create(&student)
	return c.JSON(http.StatusCreated, student)
}

func saveClass(c echo.Context) error {
	class := database.Class{}
	if err := c.Bind(&class); err != nil {
		return err
	}
	database.DB.Create(&class)
	return c.JSON(http.StatusCreated, class)
}

func searchStudentsByClasss(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.ParseUint(id, 10, 64)

	student := database.Student{ClassID: idInt}
	if err := c.Bind(&student); err != nil {
		return err
	}

	database.DB.Find(&student, id)
	return c.JSON(http.StatusOK, student)
}

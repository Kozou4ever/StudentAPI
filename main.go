package main

import (
	"StudentAPI/database"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	database.Connect()
	postgresDB, _ := database.DB.DB()
	defer postgresDB.Close()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/student", saveStudent)
	e.GET("/student/:id", getStudent)
	e.PUT("/student/:id", updateStudent)
	e.DELETE("/student/:id", deleteStudent)

	e.POST("/class", saveClass)
	e.GET("/student/:id", getClass)
	e.PUT("/student/:id", updateClass)
	e.DELETE("/student/:id", deleteClass)

	e.GET("/class/students/:id", searchStudentsByClass)

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

func getStudent(c echo.Context) error {
	id := c.Param("id")
	student := database.Student{}
	database.DB.Find(&student, id)
	return c.JSON(http.StatusOK, student)
}

func updateStudent(c echo.Context) error {
	id := c.Param("id")
	student := database.Student{}
	if err := c.Bind(&student); err != nil {
		return err
	}
	database.DB.Model(student).Where("id = ?", id).Updates(&student)
	return c.JSON(http.StatusOK, student)
}

func deleteStudent(c echo.Context) error {
	id := c.Param("id")
	student := database.Student{}
	database.DB.Delete(&student, id)
	return c.JSON(http.StatusOK, student)
}

func getClass(c echo.Context) error {
	id := c.Param("id")
	class := database.Class{}
	database.DB.Find(&class, id)
	return c.JSON(http.StatusOK, class)
}

func updateClass(c echo.Context) error {
	id := c.Param("id")
	class := database.Class{}
	if err := c.Bind(&class); err != nil {
		return err
	}
	database.DB.Model(class).Where("id = ?", id).Updates(&class)
	return c.JSON(http.StatusOK, class)
}

func deleteClass(c echo.Context) error {
	id := c.Param("id")
	class := database.Class{}
	database.DB.Delete(&class, id)
	return c.JSON(http.StatusOK, class)
}

func saveClass(c echo.Context) error {
	class := database.Class{}
	if err := c.Bind(&class); err != nil {
		return err
	}
	database.DB.Create(&class)
	return c.JSON(http.StatusCreated, class)
}

func searchStudentsByClass(c echo.Context) error {
	id := c.Param("id")

	student := database.Student{}
	if err := c.Bind(&student); err != nil {
		return err
	}

	database.DB.Delete(&student, id)
	return c.String(http.StatusOK, "User deleted!")
}

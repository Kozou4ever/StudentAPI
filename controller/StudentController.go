package controller

import (
	"StudentAPI/config"
	"StudentAPI/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateStudent(c echo.Context) error {
	student := model.Student{}
	db := config.DB()

	if err := c.Bind(&student); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	if err := db.Create(&student).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": student,
	}

	return c.JSON(http.StatusCreated, response)
}

func GetStudent(c echo.Context) error {
	id := c.Param("id")
	db := config.DB()

	var students []*model.Student

	if res := db.Find(&students, id); res.Error != nil {
		data := map[string]interface{}{
			"message": res.Error.Error(),
		}

		return c.JSON(http.StatusOK, data)
	}

	response := map[string]interface{}{
		"data": students[0],
	}

	return c.JSON(http.StatusOK, response)
}

func UpdateStudent(c echo.Context) error {
	id := c.Param("id")
	student := model.Student{}
	db := config.DB()

	if err := c.Bind(student); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	existingStudent := new(model.Student)

	if err := db.First(&existingStudent, id).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusNotFound, data)
	}

	existingStudent.StudentName = student.StudentName
	if err := db.Save(&existingStudent).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": existingStudent,
	}

	return c.JSON(http.StatusOK, response)
}

func DeleteStudent(c echo.Context) error {
	id := c.Param("id")
	student := model.Student{}
	db := config.DB()

	err := db.Delete(&student, id).Error
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "A student has been deleted.",
	}
	return c.JSON(http.StatusOK, response)
}

type BestStudent struct {
	ClassName string `json:"class_name"`
	Name      string `json:"student_name"`
	Value     int    `json:"value"`
}

func GetBestStudents(c echo.Context) error {
	db := config.DB()
	var bestStudents []BestStudent

	// Select all classes
	var classes []model.Class
	db.Find(&classes)

	// Loop through each class
	for _, class := range classes {
		var bestStudent BestStudent
		bestStudent.ClassName = class.ClassName

		// Select all marks for the current class
		var marks []model.Mark
		db.Model(&class).Association("Marks").Find(&marks)

		// Find the mark with the highest value for the current class
		maxValue := 0
		var bestStudentID uint
		for _, mark := range marks {
			if int(mark.Value) > maxValue {
				maxValue = int(mark.Value)
				bestStudentID = uint(mark.StudentID)
			}
		}

		// Get the student with the highest mark for the current class
		var student model.Student
		db.First(&student, bestStudentID)
		bestStudent.Name = student.StudentName
		bestStudent.Value = maxValue
		bestStudents = append(bestStudents, bestStudent)
	}

	return c.JSON(http.StatusOK, bestStudents)
}

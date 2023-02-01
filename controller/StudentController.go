package controller

import (
	"StudentAPI/config"
	"StudentAPI/model"
	"net/http"
	"strconv"

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

func GetBestStudentInClass(c echo.Context) error {
	db := config.DB()
	classID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		data := map[string]interface{}{
			"message": "Invalid class ID",
		}
		return c.JSON(http.StatusBadRequest, data)
	}

	var student model.Student
	var marks []model.Mark
	
	err = db.Where("class_id = ?", classID).Find(&marks).Error
	if err != nil {
		data := map[string]interface{}{
			"message": "No marks found for the class",
		}
		return c.JSON(http.StatusBadRequest, data)
	}

	var highestMark float64
	var bestStudentID uint

	for _, mark := range marks {
		if mark.Value > highestMark {
			highestMark = mark.Value
			bestStudentID = mark.StudentID
		}
	}

	err = db.Where("id = ?", bestStudentID).First(&student).Error
	if err != nil {
		data := map[string]interface{}{
			"message": "No student found with the highest mark",
		}
		return c.JSON(http.StatusBadRequest, data)
	}

	response := map[string]interface{}{
		"data": &student,
	}

	return c.JSON(http.StatusOK, response)
}

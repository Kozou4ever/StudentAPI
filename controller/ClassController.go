package controller

import (
	"StudentAPI/config"
	"StudentAPI/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func CreateClass(c echo.Context) error {
	class := model.Class{}
	db := config.DB()

	if err := c.Bind(&class); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	if err := db.Create(&class).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": class,
	}

	return c.JSON(http.StatusCreated, response)
}

func GetClass(c echo.Context) error {
	id := c.Param("id")
	db := config.DB()

	var classes []*model.Student

	if res := db.Find(&classes, id); res.Error != nil {
		data := map[string]interface{}{
			"message": res.Error.Error(),
		}

		return c.JSON(http.StatusOK, data)
	}

	response := map[string]interface{}{
		"data": classes[0],
	}

	return c.JSON(http.StatusOK, response)
}

func UpdateClass(c echo.Context) error {
	id := c.Param("id")
	class := model.Class{}
	db := config.DB()

	if err := c.Bind(class); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	existingClass := new(model.Class)

	if err := db.First(&existingClass, id).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusNotFound, data)
	}

	existingClass.ClassName = class.ClassName
	if err := db.Save(&existingClass).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": existingClass,
	}

	return c.JSON(http.StatusOK, response)
}

func DeleteClass(c echo.Context) error {
	id := c.Param("id")
	class := model.Class{}
	db := config.DB()

	err := db.Delete(&class, id).Error
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "A class has been deleted.",
	}
	return c.JSON(http.StatusOK, response)
}

func GetClassStudents(c echo.Context) error {
	classID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	db := config.DB()

	var students []model.Student
	if err := db.Table("student_classes").
		Select("students.*").
		Joins("inner join students on student_classes.student_id = students.id").
		Where("student_classes.class_id = ?", classID).
		Find(&students).
		Error; err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	response := map[string]interface{}{
		"data": students,
	}

	return c.JSON(http.StatusOK, response)
}

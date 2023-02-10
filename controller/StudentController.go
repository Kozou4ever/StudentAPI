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

func GetStudentDetails(c echo.Context) error {
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

func UpdateStudentDetails(c echo.Context) error {
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

func RankStudents(c echo.Context) error {
	db := config.DB()
	rank, _ := strconv.ParseUint(c.Param("rank"), 10, 64)

	var bestStudents []struct {
		ClassName   string
		StudentName string
		Value       int
		Rank        int
	}

	err := db.Table("(SELECT classes.class_name as class_name, students.student_name as student_name, marks.value, ROW_NUMBER() OVER (PARTITION BY classes.id ORDER BY marks.value DESC) as rank FROM marks JOIN students ON marks.student_id = students.id JOIN classes ON marks.class_id = classes.id) as BestStudent").
		Select("class_name, student_name, value, rank").
		Where("rank = ?", rank).
		Scan(&bestStudents).Error

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, bestStudents)
}

func GetTopStudentsClass(c echo.Context) error {
	return nil //todo
}

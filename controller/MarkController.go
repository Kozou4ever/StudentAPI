package controller

import (
	"StudentAPI/config"
	"StudentAPI/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateMark(c echo.Context) error {
	mark := model.Mark{}
	db := config.DB()

	if err := c.Bind(&mark); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	if err := db.Create(&mark).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": mark,
	}

	return c.JSON(http.StatusCreated, response)
}

func GetMark(c echo.Context) error {
	id := c.Param("id")
	db := config.DB()

	var marks []*model.Mark

	if res := db.Find(&marks, id); res.Error != nil {
		data := map[string]interface{}{
			"message": res.Error.Error(),
		}

		return c.JSON(http.StatusOK, data)
	}

	response := map[string]interface{}{
		"data": marks[0],
	}

	return c.JSON(http.StatusOK, response)
}

func UpdateMark(c echo.Context) error {
	id := c.Param("id")
	mark := model.Mark{}
	db := config.DB()

	if err := c.Bind(mark); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	existingMark := new(model.Mark)

	if err := db.First(&existingMark, id).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusNotFound, data)
	}

	existingMark.Value = mark.Value
	if err := db.Save(&existingMark).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": existingMark,
	}

	return c.JSON(http.StatusOK, response)
}

func DeleteMark(c echo.Context) error {
	id := c.Param("id")
	mark := model.Mark{}
	db := config.DB()

	err := db.Delete(&mark, id).Error
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "A mark has been deleted.",
	}
	return c.JSON(http.StatusOK, response)
}

package bookController

import (
	"github.com/labstack/echo/v4"
	"learn_orm/config"
	"learn_orm/models"
	"net/http"
	"strconv"
)

// get all books
func GetBooksController(c echo.Context) error {
	var books []models.Book
	if err := config.DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"books":   books,
	})
}

// get book by id
func GetBookController(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 16, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Id must be a number",
			"books":   models.Book{},
		})
	}

	var book models.Book
	if err := config.DB.First(&book, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if book.ID == 0 {
		return c.JSON(http.StatusNoContent, map[string]interface{}{
			"message": "Data not found",
			"book":    models.Book{},
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get book",
		"book":    book,
	})
}

// create book
func CreateBookController(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)

	if err := config.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"book":    book,
	})
}

// delete book by id
func DeleteBookController(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 16, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Id must be a number",
			"books":   models.Book{},
		})
	}

	var book models.Book
	if err := config.DB.Delete(&book, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var books []models.Book
	if err := config.DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete book",
		"books":   books,
	})
}

// update book by id
func UpdateBookController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 16, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Id must be a number",
			"books":   models.Book{},
		})
	}

	var book models.Book
	if err := config.DB.First(&book, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	bookParam := models.Book{}
	c.Bind(&bookParam)

	if err := config.DB.Model(models.Book{}).Where("ID = ?", id).Updates(bookParam).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update book",
		"book":    book,
	})
}

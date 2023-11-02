package handler

import (
	"employee-manager/internal/repository"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ReadEmployees(c echo.Context, db *gorm.DB) error {
	employees, err := repository.GetEmployees(db)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to retrieve employees")
	}
	return c.JSON(http.StatusOK, employees)
}

func ReadEmployeeById(c echo.Context, db *gorm.DB) error {
	id := c.Param("id")

	employee, err := repository.GetEmployeeByID(db, id)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to retrieve employee")
	}
	return c.JSON(http.StatusOK, employee)
}

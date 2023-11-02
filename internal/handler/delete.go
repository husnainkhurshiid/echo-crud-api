package handler

import (
	"employee-manager/internal/repository"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func DeleteEmployee(c echo.Context, db *gorm.DB) error {
	id := c.Param("id")

	err := repository.DeleteEmployeeById(db, id)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to retrieve employees")
	}
	return c.String(http.StatusAccepted, "Employee deleted successfully")
}

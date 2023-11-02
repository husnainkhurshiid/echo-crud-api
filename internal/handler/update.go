package handler

import (
	"employee-manager/internal/model"
	"employee-manager/internal/repository"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func UpdateEmployee(c echo.Context, db *gorm.DB) error {
	id := c.Param("id")

	updatedEmployee := new(model.Employee)

	if err := c.Bind(&updatedEmployee); err != nil {
		return c.String(http.StatusBadRequest, "Binding failed")
	}

	err := repository.UpdateEmployeeById(db, id, updatedEmployee)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to update employee data")
	}

	return c.String(http.StatusOK, "Employee updated successfully")
}

package handler

import (
	"employee-manager/internal/model"
	"employee-manager/internal/repository"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateEmployee(c echo.Context, db *gorm.DB) error {
	createEmployee := new(model.Employee)
	if err := c.Bind(createEmployee); err != nil {
		return c.String(http.StatusBadRequest, "Binding failed")
	}

	uniqueID := uuid.New()
	createEmployee.Id = uniqueID.String()[:5]

	err := repository.CreateEmployee(db, createEmployee)
	if err != nil {
		return c.String(http.StatusInternalServerError, "employee creation failed")
	} else {
		return c.String(http.StatusBadRequest, "employee created successfully")
	}
}

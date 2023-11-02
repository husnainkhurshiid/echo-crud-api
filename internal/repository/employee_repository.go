package repository

import (
	"employee-manager/internal/model"

	"gorm.io/gorm"
)

// create employee in the database
func CreateEmployee(db *gorm.DB, employee *model.Employee) error {
	if resp := db.Create(employee); resp.Error != nil {
		return resp.Error
	} else {
		if resp.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}
	}
	return nil
}

// fetch all the employees
func GetEmployees(db *gorm.DB) ([]model.Employee, error) {
	var employees []model.Employee

	if err := db.Find(&employees).Error; err != nil {
		return nil, err
	}

	return employees, nil
}

// fetch employee by id
func GetEmployeeByID(db *gorm.DB, id string) (*model.Employee, error) {
	var employee model.Employee

	if resp := db.First(&employee, "id = ?", id); resp.Error != nil {
		return &model.Employee{}, resp.Error
	} else {
		if resp.RowsAffected == 0 {
			return nil, gorm.ErrRecordNotFound
		}
	}

	return &employee, nil
}

// delete employee by id
func DeleteEmployeeById(db *gorm.DB, id string) error {
	var employee model.Employee

	if resp := db.Delete(&employee, "id = ?", id); resp.Error != nil {
		return resp.Error
	} else {
		if resp.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}
	}
	return nil
}

func UpdateEmployeeById(db *gorm.DB, id string, updatedEmployee *model.Employee) error {
	var employee model.Employee

	if resp := db.Model(&employee).Where("id = ?", id).Updates(updatedEmployee); resp.Error != nil {
		return resp.Error
	} else {
		if resp.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}
	}
	return nil
}

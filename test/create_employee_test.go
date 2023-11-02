package test

import (
	"employee-manager/internal/handler"
	"employee-manager/internal/model"
	db "employee-manager/pkg/database"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func TestIntegrationCreateEmployee(t *testing.T) {
	conn, err := db.OpenDbConnectionTest()
	if err != nil {
		t.Fatalf("Database connection error %v", err)
	}
	defer db.CloseDbConnection(conn)

	e := setupEchoApp(conn)

	employee := &model.Employee{
		Name:       "John Doe",
		Position:   "Go Associate",
		Department: "Backend",
		Email:      "john@doe.com",
	}

	uniqueID := uuid.New().String()[:5]
	employee.Id = uniqueID
	EmployeeID = uniqueID

	req := httptest.NewRequest(http.MethodPost, "/create", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("employee", employee)

	if err := handler.CreateEmployee(c, conn); err != nil {
		t.Errorf("Error creating employee: %v", err)
	}

	if rec.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, rec.Code)
	}

}

func setupEchoApp(db *gorm.DB) *echo.Echo {
	router := echo.New()

	router.POST("/create", func(c echo.Context) error {
		return handler.CreateEmployee(c, db)
	})

	router.GET("/read", func(c echo.Context) error {
		return handler.ReadEmployees(c, db)
	})

	router.GET("/update/:id", func(c echo.Context) error {
		return handler.UpdateEmployee(c, db)
	})

	router.DELETE("/delete/:id", func(c echo.Context) error {
		return handler.DeleteEmployee(c, db)
	})

	return router
}

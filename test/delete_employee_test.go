package test

import (
	"employee-manager/internal/handler"
	db "employee-manager/pkg/database"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestDeleteEmployee(t *testing.T) {
	conn, err := db.OpenDbConnectionTest()
	if err != nil {
		t.Fatalf("Database connection error %v", err)
	}
	defer db.CloseDbConnection(conn)

	e := echo.New()

	req := httptest.NewRequest(http.MethodDelete, "/delete/"+EmployeeID, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err = handler.DeleteEmployee(c, conn)
	if err != nil {
		fmt.Printf("Employee deletion error : %v", err)
	}
	assert.Equal(t, http.StatusAccepted, rec.Code)
	assert.NoError(t, err)
}

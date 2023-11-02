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

func TestReadEmployees(t *testing.T) {
	conn, err := db.OpenDbConnectionTest()
	if err != nil {
		t.Fatalf("Database connection error %v", err)
	}
	defer db.CloseDbConnection(conn)

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/employees", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err = handler.ReadEmployees(c, conn)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NoError(t, err)
}

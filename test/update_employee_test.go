package test

import (
	"bytes"
	"employee-manager/internal/handler"
	db "employee-manager/pkg/database"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestUpdateEmployee(t *testing.T) {

	conn, err := db.OpenDbConnectionTest()
	if err != nil {
		t.Fatalf("Database connection error %v", err)
	}
	defer db.CloseDbConnection(conn)

	e := echo.New()

	reqBody := []byte(`{"Name":"Husnain", "Position":"Golang Dev", "Department":"Backend", "Email":"husnain@updated.com"}`)
	req := httptest.NewRequest(http.MethodPut, "/update/"+EmployeeID, bytes.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err = handler.UpdateEmployee(c, conn)
	if err != nil {
		fmt.Printf("Error in updating employee : %v", err)
	}
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NoError(t, err)
}

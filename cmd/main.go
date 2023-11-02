package main

import (
	"employee-manager/internal/handler"
	db "employee-manager/pkg/database"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {

	conn, err := db.OpenDbConnection()
	if err != nil {
		fmt.Printf("Database connection error %v", err)
		return
	}
	defer db.CloseDbConnection(conn)

	router := echo.New()
	router.Debug = true

	router.POST("/create", func(c echo.Context) error {
		return handler.CreateEmployee(c, conn)
	})

	router.GET("/read", func(c echo.Context) error {
		return handler.ReadEmployees(c, conn)
	})

	router.GET("/read/:id", func(c echo.Context) error {
		return handler.ReadEmployeeById(c, conn)
	})

	router.PATCH("/update/:id", func(c echo.Context) error {
		return handler.UpdateEmployee(c, conn)
	})

	router.DELETE("/delete/:id", func(c echo.Context) error {
		return handler.DeleteEmployee(c, conn)
	})

	// address := os.Getenv("PORT")
	if err := router.Start(":8080"); err != nil {
		fmt.Printf("Error starting the server: %v", err)
	}

}

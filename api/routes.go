package api

import (
	"github.com/labstack/echo"
)

func createRoutes(e *echo.Echo) {
	createApplicationRoutes(e)
}

func createApplicationRoutes(e *echo.Echo) {
	e.POST("/application", createApplicationHandler)
	e.GET("/application/:id", getApplicationHandler)
	e.PUT("/application/:id", updateApplicationHandler)
	e.DELETE("/application/:id", deleteApplicationHandler)
}

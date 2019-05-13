package api

import (
	"net/http"

	"github.com/labstack/echo"
)

func createApplicationHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func getApplicationHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func updateApplicationHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func deleteApplicationHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func createWorkflowHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func getWorkflowHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func updateWorkflowHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func deleteWorkflowHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

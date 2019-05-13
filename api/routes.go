package api

import (
	"github.com/labstack/echo"
)

func createRoutes(e *echo.Echo) {
	createApplicationRoutes(e)
}

func createApplicationRoutes(e *echo.Echo) {
	e.POST("/application", createApplicationHandler)
	e.GET("/application/:application_id", getApplicationHandler)
	e.PUT("/application/:application_id", updateApplicationHandler)
	e.DELETE("/application/:application_id", deleteApplicationHandler)
}

func createWorkflowRoutes(e *echo.Echo) {
	e.POST("/application/:id/workflow", createWorkflowHandler)
	e.GET("/application/:application_id/workflow/:workflow_id", getWorkflowHandler)
	e.PUT("/application/:application_id/workflow/:workflow_id", updateWorkflowHandler)
	e.DELETE("/application/:application_id/workflow/:workflow_id", deleteWorkflowHandler)
}

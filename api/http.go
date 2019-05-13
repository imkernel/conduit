package api

import (
	"github.com/labstack/echo"
)

// StartServer starts the API server
func StartServer() {
	e := echo.New()

	e.Logger.Fatal(e.Start(":1323"))
}

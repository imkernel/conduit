package api

import "github.com/gin-gonic/gin"

// StartServer starts the API server
func StartServer() {
	r := gin.Default()
	createRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}

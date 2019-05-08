package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func createApplicationHandler(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, "Bad request")
		return
	}

	c.JSON(http.StatusOK, string(data))
}

package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func EventGet(c *gin.Context) {

	c.JSON(http.StatusOK, map[string]string{
		"key": "value",
	})
}

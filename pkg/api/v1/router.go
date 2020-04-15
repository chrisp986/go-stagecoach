package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router(r *gin.Engine) *gin.RouterGroup {

	router := r.Group("api/v1")
	{
		router.GET("/event/:id", GetEvent)
		router.POST("/event", GetEvent)
	}

	r.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})

	return router
}

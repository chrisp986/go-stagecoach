package v1

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) *gin.RouterGroup {

	router := r.Group("api/v1")
	{
		router.GET("/ping", EventGet)
	}

	return router
}

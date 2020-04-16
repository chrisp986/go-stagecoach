package v1

import (
	"github.com/chrisp986/go-stagecoach/pkg/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ApplyRoutes(r *gin.Engine) *gin.RouterGroup {

	apiv1 := r.Group("api/v1")
	{
		apiv1.GET("/event/:id", GetEvent)
		apiv1.POST("/add", dto.GetDTO)
	}

	r.NoRoute(func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusNotFound, "404 - not found")
	})

	return apiv1
}

package v1

import (
	"net/http"

	"github.com/chrisp986/go-stagecoach/pkg/dto"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.Engine) (*gin.RouterGroup, *gin.RouterGroup) {

	// r.Delims("{[{", "}]}")
	// r.SetFuncMap(template.FuncMap{
	// 	"GetInputFromForm": app.GetInputFromForm,
	// })

	r.LoadHTMLGlob("web/template/*.html")

	apiv1 := r.Group("api/v1")
	{
		apiv1.GET("/event/:id", GetEvent)
		apiv1.POST("/add", dto.GetDTO)
	}

	webservice := r.Group("/")
	{
		webservice.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"title": "StageCoach Index",
			})
		})
		webservice.GET("/login", func(c *gin.Context) {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"title": "StageCoach Login",
			})
		})
	}

	r.NoRoute(func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusNotFound, "404 - not found")
	})

	return apiv1, webservice
}

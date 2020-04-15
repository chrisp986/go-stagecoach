package v1

import (
	"github.com/chrisp986/go-stagecoach/pkg/dao"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetEvent(c *gin.Context) {

	e := dao.EventDAO{}
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	if event, err := e.Get(uint32(id)); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Printf("Error on Get request for ID: %d with code: %v", id, err)
	} else {
		c.JSON(http.StatusOK, event)
	}
}

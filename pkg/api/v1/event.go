package v1

import (
	"github.com/chrisp986/go-stagecoach/pkg/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func EventGet(c *gin.Context) {

	se := service.Event{}
	event, err := se.GetOne(3)
	if err != nil {
		log.Printf("Error in event.GetOne: %v", err)
	}
	//fmt.Println(event)

	c.JSON(http.StatusOK, map[string]string{
		"router": event.UniqueID,
	})
}

//c.JSON(http.StatusOK, map[string]string{
//"router": "v1",
//})

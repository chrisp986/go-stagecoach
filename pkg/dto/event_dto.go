package dto

import (
	"fmt"
	"github.com/chrisp986/go-stagecoach/pkg/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type EventDTO struct{}

//TODO Added client authentication and check that before accepting a DTO

func GetDTO(c *gin.Context) {

	var e model.Event

	if err := c.ShouldBindJSON(&e); err == nil {
		c.AbortWithStatusJSON(http.StatusNotFound, "404 - No event found, check ID")
		log.Printf("Error on GetDTO request for  with code: %v", err)
	} else {
		c.JSON(http.StatusCreated, fmt.Sprintf("201 - New request received"))
	}
	log.Println("DTO: ", e)
}

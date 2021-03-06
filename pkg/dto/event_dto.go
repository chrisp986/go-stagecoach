package dto

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chrisp986/go-stagecoach/pkg/model"
	"github.com/chrisp986/go-stagecoach/pkg/service"
	"github.com/gin-gonic/gin"
)

//TODO Added client authentication and check that before accepting a DTO

func GetDTO(c *gin.Context) {

	var e model.Event
	if err := c.ShouldBindJSON(&e); err != nil {
		c.AbortWithStatusJSON(http.StatusNoContent, "204 - No Content")
		log.Printf("Error on GetDTO request for  with code: %v", err)
	} else {

		eventAdded := service.EventService(e)

		if !eventAdded {
			c.AbortWithStatusJSON(http.StatusBadRequest, "400 - Error in request, no event created")
			log.Printf("Error in dto.event_dto.GetDTO: %v", err)
		} else {
			c.JSON(http.StatusCreated, fmt.Sprintf("201 - New request received"))
			//log.Printf("New Event has ID: %d", id)
			//service.updateSendDate(id)
		}
	}
}

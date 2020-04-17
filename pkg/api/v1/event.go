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

	if event, err := e.GetDAO(uint32(id)); err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, "404 - No event found, check ID")
		log.Printf("Error on GetDAO request for ID: %d with code: %v", id, err)
	} else {
		c.JSON(http.StatusOK, event)
	}
}

//func AddEvent(c *gin.Context) {
//
//}

//var e model.Event
//
//e.Sender = c.Query("sender")
//e.Receiver = c.Query("receiver")
//e.Event = c.Query("event")
//template, _ := strconv.ParseInt(c.Query("template"), 10, 16)
//e.Template = uint16(template)
//e.Subtitle = c.Query("subtitle")
//e.Body = c.Query("body")
//
//c.JSON(http.StatusOK, fmt.Sprintf(" %v", c.Query("id")))

//"id": e.ID,
//"unique_id": "blala",
//"sender": e.Sender,
//"receiver": e.Receiver,
//"event": e.Event,
//"subtitle": e.Subtitle,
//"body": e.Body,
//"template": e.Template,
//"created": e.Created,
//"sent": e.Sent,

//func AddEvent (rw http.ResponseWriter, req *http.Request) {
//
//	d := json.NewDecoder(req.Body)
//	d.DisallowUnknownFields() // error if user sends extra data
//
//	//// anonymous struct type: handy for one-time use
//	//t := struct {
//	//	Test *string `json:"test"` // pointer to string, so we can test for field absence
//	//}{}
//
//	e := model.Event{}
//
//	err := d.Decode(&e)
//	if err != nil {
//		// bad JSON or unrecognized json field
//		http.Error(rw, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	//if e.Sender == "" {
//	//	http.Error(rw, "missing field 'test' from JSON object", http.StatusBadRequest)
//	//	return
//	//}
//
//	// optional check
//	if d.More() {
//		http.Error(rw, "extraneous data after JSON object", http.StatusBadRequest)
//		return
//	}
//
//	// got all fields we expected: no more, no less
//
//	log.Println(e)
//}

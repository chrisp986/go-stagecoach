package service

import (
	cryptoRand "crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"log"
	mathRand "math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/chrisp986/go-stagecoach/pkg/db"
	"github.com/chrisp986/go-stagecoach/pkg/model"
)

//When we have data nicely loaded into our models, we can perform additional logic
//to process the data before we serve it, that’s where Services come into play.
//This extra logic can be, for example filtering, aggregating, modifying structure or validating data.
//On top of that it allows us to separate db queries from business logic, which makes the code much cleaner,
//easier to maintain and most
//importantly (for me) easier to test (More on that later). So, let’s look at the code:

//Do extra logic with the data we got from the query or api

//EventService is taking a Event Struct and tries to send a mail
func EventService(e model.Event) bool {

	log.Println("--------------------------------")
	eventAdded, id, err := addEvent(e)
	if !eventAdded || err != nil {
		log.Printf("Error in service.EventService(): %v", err)
		return false
	}
	//TODO create retry function
	log.Printf("New Event created with ID: %d", id)
	log.Printf("Sender: %s   Receiver: %s", e.Sender, e.Receiver)
	log.Printf("Template: %s", e.Template)
	eventSent, err := updateSendDate(id)
	if !eventSent || err != nil {
		log.Printf("Error in updateSendDate() %v", err)
		return false
	}
	checkMailSentInEventBuffer(id)
	return true
}

// Add creates a new Event
func addEvent(e model.Event) (bool, uint32, error) {

	sqliteDB := db.GetDB()
	var newEvent model.Event
	errEmail := errors.New("no valid Email address")
	errEmpty := errors.New("field is empty")

	c1 := make(chan string)

	go func() {
		uid := createUID(16)
		c1 <- uid
	}()

	if !validateNotEmpty(e.Sender) || !validateNotEmpty(e.Receiver) || !validateNotEmpty(e.Subtitle) || !validateNotEmpty(e.Body) || !validateNotEmpty(e.Template) {
		return false, 0, errEmpty
	}

	if !validateEmail(e.Sender) {
		log.Printf("%s is no valid Email", e.Sender)
		return false, 0, errEmail
	}

	if !validateEmail(e.Receiver) {
		log.Printf("%s is no valid Email", e.Receiver)
		return false, 0, errEmail
	}

	newEvent.UniqueID = <-c1
	newEvent.Sender = e.Sender
	newEvent.Receiver = e.Receiver
	newEvent.Subtitle = e.Subtitle
	newEvent.Body = e.Body
	newEvent.Template = e.Template

	stmt, err := sqliteDB.Prepare("INSERT INTO event_buffer(unique_id, sender, receiver, template, subtitle, body) VALUES(?, ?, ?, ?, ?, ?)")

	if err != nil {
		log.Printf("Error in Prepare addEvent() %v", err)
		return false, 0, err
	}

	res, err := stmt.Exec(
		newEvent.UniqueID,
		newEvent.Sender,
		newEvent.Receiver,
		newEvent.Subtitle,
		newEvent.Body,
		newEvent.Template)

	if err != nil {
		log.Printf("Error on Exec in addEvent(): %v", err)
		return false, 0, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		log.Printf("Error on LastInsertId() in event addEvent(): %v", err)
		return false, 0, err
	}
	return true, uint32(lastID), err
}

func checkMailSentInEventBuffer(id uint32) {

	sqliteDB := db.GetDB()
	var e model.Event

	err := sqliteDB.Get(&e, "SELECT id, unique_id, sender, receiver, template, created FROM event_buffer WHERE id= ? LIMIT 1", id)
	if err != nil {
		log.Printf("Error in checkMailSentInEventBuffer() %v", err)
	}
	dateCreated, err := time.Parse(time.RFC3339, e.Created)
	if err != nil {
		log.Printf("Error in time.Parse %v", dateCreated)
	}
}

func updateSendDate(id uint32) (bool, error) {

	sqliteDB := db.GetDB()

	stmt, err := sqliteDB.Prepare("UPDATE event_buffer SET sent_date = (STRFTIME('%d-%m-%Y  %H:%M:%f', 'NOW'," +
		"'localtime')), sent = ? WHERE id = ?")

	if err != nil {
		log.Printf("Error in Prepare updateSendDate() %v", err)
		return false, err
	}
	res, err := stmt.Exec(1, id)

	if err != nil {
		log.Printf("Error on Exec in updateSendDate(): %v", err)
		return false, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error on RowsAffected() in event updateSendDate(): %v", err)
		return false, err
	}
	if rowsAffected != 0 {
		log.Println("Mail has been sent!")
		return true, err
	}
	return false, err
}

//createUID creates a unique ID based in the crypto/rand function, parameter is the size of the byte,
//return is the uid as string
func createUID(n int) string {
	b := make([]byte, n)
	_, err := cryptoRand.Read(b[:])
	if err != nil {
		log.Println(err)
	}

	mathRand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
	uid := hex.EncodeToString(b[:])

	return uid
}

func validateEmail(email string) bool {
	var rxEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if len(email) > 254 || !rxEmail.MatchString(email) {
		return false
	}
	return true
}

func validateNotEmpty(value string) bool {
	if strings.TrimSpace(value) == "" {
		return false
	}
	return true
}

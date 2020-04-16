package service

import (
	cryptoRand "crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"github.com/chrisp986/go-stagecoach/pkg/db"
	"github.com/chrisp986/go-stagecoach/pkg/model"
	"log"
	mathRand "math/rand"
	"net/http"
)

//When we have data nicely loaded into our models, we can perform additional logic
//to process the data before we serve it, that’s where Services come into play.
//This extra logic can be, for example filtering, aggregating, modifying structure or validating data.
//On top of that it allows us to separate db queries from business logic, which makes the code much cleaner,
//easier to maintain and most
//importantly (for me) easier to test (More on that later). So, let’s look at the code:

//Do extra logic with the data we got from the query or api

type Event []model.Event

type Adder interface {
	Add(event model.Event) (id uint32, err error)
}

type Buffer struct {
	Events []model.Event
}

func New() *Buffer {
	return &Buffer{Events: []model.Event{}}
}

func (e *Buffer) AddEvent(event model.Event) {
	e.Events = append(e.Events, event)
}

func NewEvent(event Adder) http.HandlerFunc {

	c1 := make(chan string)

	go func() {
		uid := createUID(16)
		c1 <- uid
	}()

	return func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}

		json.NewDecoder(r.Body).Decode(&request)

		event.Add(model.Event{
			UniqueID: <-c1,
			Sender:   "testsender@uhf.com",
			Receiver: "blkjsdjf@ijdsa.com",
			Event:    "cr",
			Subtitle: "subtitleniuenf",
			Body:     "bodyiwejfw",
			Template: 0,
		})
	}
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

// Add creates a new Event
func (e Event) Add() (id uint32, err error) {

	var event model.Event
	sqliteDB := db.GetDB()

	c1 := make(chan string)

	go func() {
		uid := createUID(16)
		c1 <- uid
	}()

	event.UniqueID = <-c1
	event.Sender = "test432@abs.com"
	event.Receiver = "testreceiver@test123.com"
	event.Event = "cr"
	event.Subtitle = "subtitle"
	event.Body = "testbody"
	event.Template = 1

	log.Println("Data to insert into event_buffer:", event)

	stmt, err := sqliteDB.Prepare("INSERT INTO event_buffer(unique_id, sender, receiver, event, subtitle, body, " +
		"template) VALUES(?, ?, ?, ?, ?, ?, ?)")

	if err != nil {
		log.Printf("Error in Prepare event.Add %v", err)
	}

	res, err := stmt.Exec(
		event.UniqueID,
		event.Sender,
		event.Receiver,
		event.Event,
		event.Subtitle,
		event.Body,
		event.Template)

	if err != nil {
		log.Printf("Error on Exec in event Add(): %v", err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		log.Printf("Error on LastInsertId() in event Add(): %v", err)
	}
	log.Printf("Event inserted with ID: %d", lastId)

	return uint32(lastId), err
}

//// Get just retrieves using User DAO, here can be additional logic for processing data retrieved by DAOs
//func (e Event) Get(id uint32) (*model.Event, error) {
//
//	var event model.Event
//	sqliteDB := db.GetDB()
//
//	err := sqliteDB.Get(&event, "SELECT * FROM event_buffer WHERE id = ?", id)
//
//	return &event, err
//}

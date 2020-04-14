package service

import (
	"github.com/chrisp986/go-stagecoach/pkg/model"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

//When we have data nicely loaded into our models, we can perform additional logic
//to process the data before we serve it, that’s where Services come into play.
//This extra logic can be, for example filtering, aggregating, modifying structure or validating data.
//On top of that it allows us to separate database queries from business logic, which makes the code much cleaner,
//easier to maintain and most
//importantly (for me) easier to test (More on that later). So, let’s look at the code:

//Do extra logic with the data we got from the query or api

type Event []model.Event

// Get just retrieves user using User DAO, here can be additional logic for processing data retrieved by DAOs
func (e Event) GetOne(sqliteDB *sqlx.DB, id uint32) (*model.Event, error) {

	var event model.Event
	err := sqliteDB.Get(&event, "SELECT * FROM event_buffer WHERE id = ?", id)

	//err := sqliteDB.QueryRowx("SELECT * FROM event_buffer WHERE id=? LIMIT 1").StructScan(&event)
	return &event, err
}

func (e Event) Add(sqliteDB *sqlx.DB) error {

	var event model.Event
	t := time.Now().Unix()

	event.UUID = uint32(t)
	event.Sender = "test432@abs.com"
	event.Receiver = "testreceiver@test123.com"
	event.Event = "cr"
	event.Subtitle = "subtitle"
	event.Body = "testbody"
	event.Template = 1

	log.Println("Data to insert into event_buffer:", event)

	stmt, err := sqliteDB.Prepare("INSERT INTO event_buffer(uuid, sender, receiver, event, subtitle, body, " +
		"template) VALUES(?, ?, ?, ?, ?, ?, ?)")

	if err != nil {
		log.Printf("Error in Prepare event.Add %v", err)
	}

	res, err := stmt.Exec(
		event.UUID,
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

	return err
}

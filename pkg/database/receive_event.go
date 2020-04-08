package database

import (
	"github.com/jmoiron/sqlx"
	"log"
)

//TODO add method and interface to call the functions interchangeably
//TODO write API to call the functions from outside of the application

func AddEvent(sqliteDB *sqlx.DB, e EventBuffer) {

	result, err := sqliteDB.Exec(`INSERT INTO event_buffer(sender, receiver, event, subtitle, body, template) 
									VALUES(?, ?, ?, ?, ?, ?)`,
		e.Sender,
		e.Receiver,
		e.Event,
		e.Subtitle,
		e.Body,
		e.Template)

	if err != nil {
		log.Printf("Error during Exec in AddEvent(): %v", err)
	}

	uuid, _ := result.LastInsertId()
	log.Printf("New Event has been added with ID: %d", uuid)
}

func AddMailAddress(sqliteDB *sqlx.DB, m MailAddress) {

	result, err := sqliteDB.Exec(`INSERT INTO mail_address(mail_address, first_name, name, status) 
									VALUES(?, ?, ?, ?)`,
		m.MailAddress,
		m.FirstName,
		m.Name,
		m.Status)

	if err != nil {
		log.Printf("Error during Exec in AddMailAddress(): %v", err)
	}

	id, _ := result.LastInsertId()
	log.Printf("New Mail Address has been added with ID: %d", id)
}

func AddMsgTemplate(sqliteDB *sqlx.DB, m MsgTemplate) {

	result, err := sqliteDB.Exec(`INSERT INTO msg_template(msg_subtitle, msg_body) 
									VALUES(?, ?)`,
		m.MsgSubtitle,
		m.MsgBody)

	if err != nil {
		log.Printf("Error during Exec in AddMsgTemplate(): %v", err)
	}

	id, _ := result.LastInsertId()
	log.Printf("New Message Template has been added with ID: %d", id)
}

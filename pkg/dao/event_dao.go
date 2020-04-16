package dao

import (
	"github.com/chrisp986/go-stagecoach/pkg/db"
	"github.com/chrisp986/go-stagecoach/pkg/model"
)

// EventDAO persists event data in database
type EventDAO struct{}

//Creates an new EventDAO
func NewEventDAO() *EventDAO {
	return &EventDAO{}
}

//GetDAO() queries the database and returns an event struct if and event with that ID is available
func (dao *EventDAO) GetDAO(id uint32) (*model.Event, error) {

	sqliteDB := db.GetDB()
	var e model.Event

	err := sqliteDB.Get(&e, "SELECT * FROM event_buffer WHERE id = ?", id)

	return &e, err
}

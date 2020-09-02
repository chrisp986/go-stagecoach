//MIT License
//
//Copyright (c) 2020 Christian Peters
//
//Permission is hereby granted, free of charge, to any person obtaining a copy
//of this software and associated documentation files (the "Software"), to deal
//in the Software without restriction, including without limitation the rights
//to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//copies of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in all
//copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//SOFTWARE.

package main

import (
	"log"

	"github.com/chrisp986/go-stagecoach/init/sqlite"
	"github.com/chrisp986/go-stagecoach/pkg/api"
	"github.com/chrisp986/go-stagecoach/pkg/db"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var sqliteDB *sqlx.DB
var initDone bool
var msg string

func init() {
	log.Println("->> Stagecoach v0.2.31 <<-")
	log.Println("-----------------------")

	msg, initDone = sqlite.InitiateDatabase()
	log.Println(msg)

	db.ConnectDB() // Connect to SQLite DB
	log.Println("Connection to db established.")

	if initDone {
		log.Println("Initialization completed. Starting Server on Port :8080")
		api.RunServer() // Start Gin Server
	}
}

func main() {

	sqliteDB = db.GetDB()
	defer sqliteDB.Close() //Close connection when main() stops

	//se := service.Event{}
	//err := se.Add()
	//if err != nil {
	//	log.Printf("Error in event.Add(): %v", err)
	//}

	//model, err := se.GetDAO(sqliteDB, 3)
	//if err != nil {
	//	log.Printf("Error in event.GetDAO(): %v", err)
	//}

	//es := service.Event{}
	//model, err := es.GetDAO(sqliteDB, 1)
	//log.Println(model)

	//eb := db.EventBuffer{
	//	UUID:     123456,
	//	Sender:   3,
	//	Receiver: 12,
	//	Event:    0,
	//	Subtitle: "testsub",
	//	Body:     "testbody",
	//	Template: 1,
	//}
	//
	//eb.AddEvent(sqliteDB)
	//
	//ma := db.MailAddress{
	//	MailAddress: "test@test.com",
	//	FirstName:   "Christian",
	//	Name:        "Peters",
	//	Status:      0,
	//}
	//ma.AddMailAddress(sqliteDB)
	//
	//mt := db.MsgTemplate{
	//	MsgSubtitle: "testsubtitle",
	//	MsgBody:     "testbody",
	//}
	//mt.AddMsgTemplate(sqliteDB)

}

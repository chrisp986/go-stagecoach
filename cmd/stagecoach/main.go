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
	"fmt"
	"github.com/chrisp986/go-stagecoach/init/sqlite"
	"github.com/chrisp986/go-stagecoach/pkg/database"
	"github.com/jmoiron/sqlx"
	"log"
	"path/filepath"
)

func main() {
	fmt.Println("->> Stagecoach v0.1")
	fmt.Println("-------------------")

	msg := sqlite.InitiateDatabase()
	log.Println(msg)

	sqliteDB, err := sqlx.Connect("sqlite3", filepath.Join("internal", "sqlitedb", "sqlite_database.db"))
	if err != nil {
		log.Fatalf("Connection to database %v", err)
	}
	defer sqliteDB.Close()

	log.Println("Connection to database established")

	eb := database.EventBuffer{
		UUID:     123456,
		Sender:   3,
		Receiver: 12,
		Event:    0,
		Subtitle: "testsub",
		Body:     "testbody",
		Template: 1,
	}

	eb.AddEvent(sqliteDB)

	ma := database.MailAddress{
		MailAddress: "test@test.com",
		FirstName:   "Christian",
		Name:        "Peters",
		Status:      0,
	}
	ma.AddMailAddress(sqliteDB)

	mt := database.MsgTemplate{
		MsgSubtitle: "testsubtitle",
		MsgBody:     "testbody",
	}
	mt.AddMsgTemplate(sqliteDB)

}

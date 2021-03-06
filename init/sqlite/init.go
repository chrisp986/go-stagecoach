package sqlite

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var sqlitePath = filepath.Join("internal", "sqlitedb", "sqlite_database.db")

func InitiateDatabase() (msg string, initDone bool) {

	log.Println("Start initializing.")
	var sqliteDB *sqlx.DB
	initDone = false

	if !fileExist(sqlitePath) {
		createSQLiteDB()
		sqliteDB, _ = sqlx.Open("sqlite3", sqlitePath)
		defer sqliteDB.Close()

		tablesCreated := createSQLiteTable(createTables)

		if tablesCreated {
			return "SQLite db created, initialization done.", true
		}
	}
	return "SQLite db already exists, skipping creation.", true
}

//Check if the db already exists, will not check for tables
func fileExist(name string) bool {
	_, err := os.Stat(name)
	if err != nil {
		return false
	}
	return true
}

//Creates the SQLite db
func createSQLiteDB() (msg string, dbCreated bool) {
	dbCreated = false

	log.Println("Initialize SQLite db.")

	dbDirExist := createDBPath()

	if dbDirExist {
		file, err := os.Create(sqlitePath)
		if err != nil {
			log.Printf("Error while creating SQLite db %v.", err)
		}
		defer file.Close()

		return "SQLite directory and db created.", true
	}
	return "SQLite directory and db could not be created.", false
}

func createDBPath() bool {

	dbPath := filepath.Join("internal", "sqlitedb")
	err := os.MkdirAll(dbPath, os.ModePerm)
	if err != nil {
		log.Printf("Error while creating SQLite directory %v", dbPath)
		return false
	}
	log.Println("Directory for SQLite db created.")
	return true
}

//Creates the tables in the SQLite
func createSQLiteTable(query string) bool {

	db, err := sqlx.Open("sqlite3", sqlitePath)
	if err != nil {
		log.Printf("Error while opening SQLiteDB in createSQLiteTable(): %v.", err)
		return false
	}

	stmts := strings.Split(query, ";\n")
	if len(strings.Trim(stmts[len(stmts)-1], " \n\t\r")) == 0 {
		stmts = stmts[:len(stmts)-1]
	}
	for _, s := range stmts {
		_, err := db.Exec(s)
		if err != nil {
			log.Printf("Error while creating table %v.", err)
			return false
		}
	}
	return true
}

//DATA FOR TESTING
//INSERT INTO event_buffer(sender, receiver, event, subtitle, body, sent) VALUES(1, 1, "cr", "testsubtitle", "testbody", "2020-10-07 08:23:19");
//INSERT INTO mail_from(fmail_address, first_name, name) VALUES ("cpeters986@gmail.com", "Christian", "Peters");
//INSERT INTO mail_to(tmail_address, first_name, name) VALUES ("d4m1en@gmail.com", "Christian", "Peters");
//SELECT event_buffer.uuid, event_buffer.subtitle, event_buffer.body, mail_address.mail_address, mail_address.first_name, mail_address.status FROM event_buffer INNER JOIN mail_address ON event_buffer.sender=mail_address.id WHERE mail_address.status='3';

var createTables = `CREATE TABLE IF NOT EXISTS event_buffer(
		  id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
          unique_id TEXT NOT NULL UNIQUE,
		  sender TEXT NOT NULL DEFAULT 'sender_default',
		  receiver TEXT NOT NULL DEFAULT 'receiver_default',
		  template TEXT NOT NULL DEFAULT 'template_default',
		  subtitle TEXT NOT NULL DEFAULT 'subtitle_default',
		  body TEXT NOT NULL DEFAULT 'body_default',
		  created DATETIME NOT NULL DEFAULT (STRFTIME('%d-%m-%Y  %H:%M:%f', 'NOW','localtime')),
		  sent_date DATETIME NOT NULL DEFAULT '01-01-1970  00:00:00.000',
		  sent INTEGER NOT NULL DEFAULT 0
		);

		CREATE TABLE IF NOT EXISTS mail_address(
			mail_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			mail_address TEXT NOT NULL DEFAULT 'default@default.com',
			first_name TEXT NOT NULL DEFAULT 'first_name_default',
			name TEXT NOT NULL DEFAULT 'name_default',
			status INT NOT NULL DEFAULT 0,
			created DATETIME NOT NULL DEFAULT (STRFTIME('%d-%m-%Y  %H:%M:%f', 'NOW','localtime'))
		);`

//
//CREATE TABLE IF NOT EXISTS msg_template(
//msg_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
//msg_subtitle TEXT NOT NULL,
//msg_body TEXT NOT NULL,
//created DATETIME DEFAULT (STRFTIME('%d-%m-%Y  %H:%M:%f', 'NOW','localtime'))
//);`

//Status 1 = only sender, 2 = only receiver, 0 = both
//FOREIGN KEY(event) REFERENCES msg_template(id)

//FOREIGN KEY(sender) REFERENCES mail_address(mail_id),
//FOREIGN KEY(receiver) REFERENCES mail_address(mail_id),
//FOREIGN KEY(event) REFERENCES msg_template(msg_id)

//FOREIGN KEY(sender) REFERENCES mail_address(mail_id),
//FOREIGN KEY(receiver) REFERENCES mail_address(mail_id),
//FOREIGN KEY(event) REFERENCES msg_template(msg_id)

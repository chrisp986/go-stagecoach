package sqlite

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var sqlitePath = filepath.Join("internal", "sqlitedb", "sqlite_database.db")

func InitiateDatabase() string {
	var msg string
	var sqliteDB *sqlx.DB

	if !fileExist(sqlitePath) {
		createSQLiteDB()
		sqliteDB, _ = sqlx.Open("sqlite3", sqlitePath)
		defer sqliteDB.Close()
	} else {
		msg = "SQLite database already exists"
	}

	createSQLiteTable(createMainTable)
	msg = "SQLite database created"
	return msg
}

//Check if the database already exists, will not check for tables
func fileExist(name string) bool {
	_, err := os.Stat(name)
	if err != nil {
		return false
	}
	return true
}

//Creates the SQLite database
func createSQLiteDB() {
	log.Println("Creating SQLite-database.db")
	file, err := os.Create(sqlitePath) // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()
}

//Creates the tables in the SQLite
func createSQLiteTable(query string) {

	db, err := sqlx.Open("sqlite3", sqlitePath)
	if err != nil {
		log.Fatalf("Error while opening SQLiteDB in createSQLiteTable(): %v", err)
	}

	stmts := strings.Split(query, ";\n")
	if len(strings.Trim(stmts[len(stmts)-1], " \n\t\r")) == 0 {
		stmts = stmts[:len(stmts)-1]
	}
	for _, s := range stmts {
		_, err := db.Exec(s)
		if err != nil {
			log.Fatalf("Error while creating table %v", err)
		}
	}
	log.Println("SQLite tables created")
}

//DATA FOR TESTING
//INSERT INTO event_buffer(sender, receiver, event, subtitle, body, sent) VALUES(1, 1, "cr", "testsubtitle", "testbody", "2020-10-07 08:23:19");
//INSERT INTO mail_from(fmail_address, first_name, name) VALUES ("cpeters986@gmail.com", "Christian", "Peters");
//INSERT INTO mail_to(tmail_address, first_name, name) VALUES ("d4m1en@gmail.com", "Christian", "Peters");
//SELECT event_buffer.uuid, event_buffer.subtitle, event_buffer.body, mail_address.mail_address, mail_address.first_name, mail_address.status FROM event_buffer INNER JOIN mail_address ON event_buffer.sender=mail_address.id WHERE mail_address.status='3';

var createMainTable = `CREATE TABLE IF NOT EXISTS event_buffer(
		  uuid INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
		  sender INTEGER NOT NULL,
		  receiver INTEGER NOT NULL,
		  event INTEGER NOT NULL,
		  subtitle TEXT,
		  body TEXT,
		  template INTEGER DEFAULT 0,
		  created DATETIME DEFAULT CURRENT_TIMESTAMP,
		  sent DATETIME,
		FOREIGN KEY(sender) REFERENCES mail_address(id),
		FOREIGN KEY(receiver) REFERENCES mail_address(id),
		FOREIGN KEY(event) REFERENCES msg_template(id)
		);

		CREATE TABLE IF NOT EXISTS mail_address(
		 id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		 mail_address TEXT NOT NULL,
		 first_name TEXT,
		 name TEXT,
		 status INT NOT NULL DEFAULT 0, 
		 created DATETIME DEFAULT CURRENT_TIMESTAMP
		);

		CREATE TABLE IF NOT EXISTS msg_template(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		msg_subtitle TEXT NOT NULL,
		msg_body TEXT NOT NULL,
		created DATETIME DEFAULT CURRENT_TIMESTAMP
		);`

//Status 1 = only sender, 2 = only receiver, 0 = both

package db

import (
	"log"
	"path/filepath"

	"github.com/jmoiron/sqlx"
)

var sqliteDB *sqlx.DB

func ConnectDB() {
	var err error

	sqliteDB, err = sqlx.Connect("sqlite3", filepath.Join("internal", "sqlitedb", "sqlite_database.db?_loc=auto"))
	if err != nil {
		log.Fatalf("Couldn't establish the connection to SQLite DB %v", err)
	}
	sqliteDB.SetMaxOpenConns(5)

}

//GetDB returns the handle to the SQLite database
func GetDB() *sqlx.DB {
	return sqliteDB
}

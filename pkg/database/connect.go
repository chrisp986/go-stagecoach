package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"path/filepath"
)

//ConnectToDB initiates the connection to the database
//returns the db instance
func ConnectToDB() *sqlx.DB {

	sqliteDB, err := sqlx.Connect("sqlite3", filepath.Join("internal", "sqlitedb", "sqlite_database.db"))
	if err != nil {
		log.Fatalf("Connection to database %v", err)
	}
	defer sqliteDB.Close()

	log.Println("Connection to database established")
	return sqliteDB
}

//func GetCRPhases(db *sqlx.DB) []string {
//
//	var phases []string
//	err := db.Select(&phases, fmt.Sprintf(GetCRPhasesQuery, "phase", "change_request", "phase"))
//	if err != nil {
//		log.Printf("Error in GetCRPhases() %v\n", err)
//	}
//	//fmt.Printf("CR Phases: %s\n", phases)
//	return phases
//}

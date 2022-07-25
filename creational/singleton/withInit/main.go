package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/mattn/go-sqlite3"
)

var db *sql.DB

func init() {
	// Get a database handle.
	var err error
	db, err = sql.Open("sqlite3", os.Getenv("DBURL"))
	if err != nil {
		if sqliteError, ok := err.(sqlite3.Error); ok {
			fmt.Printf("Unhandled SQLite3 error. Code:[%s] Extended:[%s] Desc:[%s]", sqliteError.Code, sqliteError.ExtendedCode, sqliteError.Error())
			return
		}
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
}

// GetDBConnection returns the instance of database for connection pooling
func GetDBConnection() *sql.DB {
	return db
}

func main() {
	db := GetDBConnection()
	fmt.Printf("DB conn address: %p\n", &db)
	db = GetDBConnection()
	fmt.Printf("DB conn address: %p", &db)
}

package db 

import (
	"log"
	"time"
	"database/sql"
	_ "github.com/lib/pq"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName, dataSourceName string) (*Adapter) {
  // connect to the database 
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("Error while connecting to DB %s\n", err)
	}
	// test the db connection 
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error while connecting to DB %s\n", err)
	}
	return &Adapter{ db: db }
}



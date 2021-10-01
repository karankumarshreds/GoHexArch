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

func (da Adapter) CloseDBConnection() {
	err := da.db.Close()
	if err != nil {
		log.Fatalf("DB Close failure: %v", err)
	}
}

func (da Adapter) AddToHistory(answer int32, operation string)  error {
	return da.db.QueryRow(
		"INSERT INTO arith_history(operation, answer, date) VALUES($1, $2) RETURNING operation, answer, date",
		operation, answer, time.Now(),
	).Scan()
}


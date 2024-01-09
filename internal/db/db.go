package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

type DB struct {
	conn *sql.DB
}

const connStr string = "user=kagura password=sandy_777 dbname=todo-db sslmode=disable"

func NewDB() (*DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err

	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to the database!")
	return &DB{conn: db}, nil
}

func (db *DB) AddTaskFromDB() {
	
}

func (db *DB) DeleteTaskFromDB(taskID int) (bool, error){
	_, err := db.conn.Exec("DELETE FROM tasks WHERE task = $1", taskID)
	
	if err != nil {
		return false, err
	}

	return true, nil
}

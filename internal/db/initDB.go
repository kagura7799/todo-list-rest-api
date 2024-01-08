package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	const connStr = "user=kagura dbname=todo-db sslmode=disable"

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
	return db, nil
}

package data

import (
	"database/sql"
	"errors"
	"os"
	"task-manager-app/cmd/config"

	_ "github.com/mattn/go-sqlite3"
)

func InitDatabase() (*sql.DB, error) {

	dbPath := "internal/data/users.db"
	_, err := os.Stat(dbPath)
	if err != nil && os.IsNotExist(err) {
		conn, connErr := sql.Open("sqlite3", dbPath)
		if connErr != nil {
			return nil, errors.New("error initializing database")
		}
		_, createErr := conn.Exec(config.CreateTable)
		if createErr != nil {
			return nil, errors.New("error creating table")
		}
		return conn, nil
	} else if err != nil {
		return nil, errors.New("error checking if database exists")
	}

	conn, connErr := sql.Open("sqlite3", dbPath)
	if connErr != nil {
		return nil, errors.New("error opening existing database")
	}

	return conn, nil
}

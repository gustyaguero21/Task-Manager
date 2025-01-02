package config

//router params

const (
	Port = ":8080"
)

//sql queries

const (
	CreateTable = `CREATE TABLE USERS (id INTEGER PRIMARY KEY, name TEXT NOT NULL, surname TEXT NOT NULL, email TEXT NOT NULL UNIQUE, password TEXT NOT NULL);`
)

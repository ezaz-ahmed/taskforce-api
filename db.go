package main

import "database/sql"

type MySQLStorage struct {
	db *sql.DB
}

func NewMySQLStorage(config mysql.Config) *MySQLStorage {
	db, err := sql.Open("mysql", config.FormatDSN())
}

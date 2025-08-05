package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	
)

var Conn *sql.DB

func InitDB(connStr string) error {
	var err error
	Conn, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	if err := Conn.Ping(); err != nil {
		return err
	}
	return nil
}


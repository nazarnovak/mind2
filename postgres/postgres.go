package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DB struct {
	Conn *sql.DB
}

func (db *DB) Open(host, port, user, password, name string) error {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s " +
		"dbname=%s sslmode=disable", host, port, user, password, name)

	db.Conn, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	err = db.Conn.Ping()
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) Close() {
	db.Conn.Close()
}
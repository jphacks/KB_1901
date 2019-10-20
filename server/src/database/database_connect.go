package database

import (
	"fmt"
	"errors"
	_ "github.com/lib/pq"
	"github.com/gocraft/dbr"
	"../config"
)

type DB struct {
	Conn *dbr.Connection
	Sess *dbr.Session
}

func Connect( database config.DBconf ) ( *DB, error ) {
	source := fmt.Sprintf(
		"user=%v dbname=%v password=%v host=%v port=%v sslmode=disable",
		database.USER,
		database.DATABASE,
		database.PASSWORD,
		database.HOST,
		database.PORT,
	)

	conn, _ := dbr.Open("postgres", source, nil)
	err := conn.Ping()

	if err != nil {
		return nil, err
	}

	sess := conn.NewSession( nil )

	return &DB{ conn, sess }, err
}

func Disconnect(db *DB) error {
	if db == nil {
		return errors.New("database instance is nil")
	}

	err := db.Sess.Close()
	if err != nil {
		return err
	}

	err = db.Conn.Close()
	return err
}


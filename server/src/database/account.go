package database

import (
	_ "github.com/lib/pq"
	"github.com/gocraft/dbr"
)

func Account_ID( sess *dbr.Session, account string ) ( int, error ) {
	var ID int

	err :=  sess.QueryRow( `SELECT ID FROM account_info WHERE name = $1`,
		account,).Scan( &ID )

	return ID, err
}



package internalsql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect(host string) (*sql.DB, error) {
	db, err := sql.Open("mysql", host)
	if err != nil {
		log.Fatalf("error connecting to database %+v\n", err)
		return nil, err
	}
	return db, nil
}

package db

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/crypto-papers/api/config"
	_ "github.com/lib/pq"
)

var once sync.Once

// Connect will establish a connection to the specified postgres database
func Connect() (*sql.DB, error) {
	var db *sql.DB
	var err error

	conf := config.New()

	once.Do(func() {
		dbinfo := fmt.Sprintf(
			"user=%s password=%s dbname=%s sslmode=%s",
			conf.Postgres.PGUser, conf.Postgres.PGPassword, conf.Postgres.PGName, conf.Postgres.PGSSL,
		)
		db, _ = sql.Open("postgres", dbinfo)
		err = db.Ping()
	})
	return db, err
}

func ExecQuery(db *sql.DB, query string, args ...interface{}) {
	_, err := db.Exec(query, args...)
	if err != nil {
		panic(err)
	}
}

package db

import (
	"database/sql"
	"fmt"
	"strings"
	"sync"

	"github.com/crypto-papers/api/config"
	_ "github.com/lib/pq"
)

var once sync.Once
var conf = config.New()

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

// ExecQuery executes a query against the specified database
func ExecQuery(db *sql.DB, query string, args ...interface{}) {
	_, err := db.Exec(query, args...)
	if err != nil {
		panic(err)
	}
}

// LogAndQuery prints the query to the log output before running it
func LogAndQuery(db *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {
	s := stripWhiteSpace(query)
	fmt.Println(s)
	return db.Query(query, args...)
}

// LogQueryAndScan prints the query to the log output before running it and scanning for a returning id
func LogQueryAndScan(db *sql.DB, query string, args ...interface{}) (string, error) {
	fmt.Println(query)

	var scan string
	err := db.QueryRow(query, args...).Scan(&scan)
	if err != nil {
		panic(err)
	}

	return scan, nil
}

// CheckSchemaVersion compares the current database schema version to that specified in the environmental variables
func CheckSchemaVersion(db *sql.DB) {
	var regclass string
	exists := db.QueryRow("SELECT to_regclass('internal.config');")
	exists.Scan(&regclass)

	if regclass == "internal.config" {
		var schemaVersion string
		row := db.QueryRow("SELECT config_value FROM internal.config WHERE config_name = 'schema_version';")
		row.Scan(&schemaVersion)

		if schemaVersion == conf.Postgres.PGSchema {
			fmt.Println("Database up to date.")
		} else {
			fmt.Println("Updating database...")
			InitDB(db)
		}
	} else {
		fmt.Println("Initializing database...")
		InitDB(db)
	}
}

// InitDB initializes the database with schema
func InitDB(db *sql.DB) {
	switch conf.Postgres.PGSchema {
	case "1.0":
		schemaOne(db)
	}
}

func stripWhiteSpace(s string) string {
	s = strings.Replace(s, "\n", " ", -1)
	s = strings.Replace(s, "\t", "", -1)
	return s
}

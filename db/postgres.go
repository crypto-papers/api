package db

import (
	"database/sql"
	"fmt"
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

func LogAndQuery(db *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {
	fmt.Println(query)
	return db.Query(query, args...)
}

// CheckSchemaVersion compares the current database schema version to that specified in the environmental variables
func CheckSchemaVersion(db *sql.DB) {
	var regclass string
	exists := db.QueryRow("SELECT to_regclass('internal.config');")
	exists.Scan(&regclass)

	if regclass == "internal.config" {
		var config_value string
		row := db.QueryRow("SELECT config_value FROM internal.config WHERE config_name = 'schema_version';")
		row.Scan(&config_value)

		if config_value == conf.Postgres.PGSchema {
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

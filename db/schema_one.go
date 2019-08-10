package db

import (
	"database/sql"
	"fmt"
)

func schemaOne(db *sql.DB) {
	// Add extension to generate uuids
	ExecQuery(db, "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	// Create internal schema
	ExecQuery(db, "CREATE SCHEMA IF NOT EXISTS internal")

	// Create required tables
	ExecQuery(db, "CREATE TABLE IF NOT EXISTS internal.config (id SERIAL PRIMARY KEY, config_name varchar(255), config_value varchar(255));")
	ExecQuery(db, "CREATE TABLE IF NOT EXISTS public.authors (id uuid DEFAULT uuid_generate_v4 (), name varchar(255), papers uuid [], psuedonym BOOL, created_at TIMESTAMPTZ, PRIMARY KEY (id));")
	ExecQuery(db, "CREATE TABLE IF NOT EXISTS public.currencies (id uuid DEFAULT uuid_generate_v4 (), name varchar(255), ticker varchar(255), created_at TIMESTAMPTZ, PRIMARY KEY (id));")
	ExecQuery(db, "CREATE TABLE IF NOT EXISTS public.files (id uuid DEFAULT uuid_generate_v4 (), coverImage varchar(255), source varchar(255), url varchar(255), created_at TIMESTAMPTZ, PRIMARY KEY (id));")
	ExecQuery(db, "CREATE TABLE IF NOT EXISTS public.papers (id uuid DEFAULT uuid_generate_v4 (), author uuid [], currency uuid [], description TEXT, excerpt TEXT, page_num SMALLINT, title varchar(255), created_at TIMESTAMPTZ, PRIMARY KEY (id));")

	// Insert default values
	ExecQuery(db, fmt.Sprintf("SET timezone = '%s';", conf.Base.Timezone))
	ExecQuery(db, fmt.Sprintf("INSERT INTO internal.config (config_name, config_value) VALUES ('schema_version', %s);", conf.Postgres.PGSchema))
}

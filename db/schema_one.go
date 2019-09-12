package db

import (
	"database/sql"
	"fmt"
)

func schemaOne(db *sql.DB) {
	// Add extension to generate uuids
	ExecQuery(db, "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	// Create internal schema
	ExecQuery(db, "CREATE SCHEMA IF NOT EXISTS internal;")

	// Create required tables
	ExecQuery(db, "CREATE TABLE IF NOT EXISTS internal.config (id SERIAL, config_name VARCHAR(255), config_value VARCHAR(255), PRIMARY KEY (id));")
	ExecQuery(db, "CREATE TABLE IF NOT EXISTS public.assets (id uuid DEFAULT uuid_generate_v4 (), name VARCHAR(255), ticker VARCHAR(255), created_at TIMESTAMPTZ, PRIMARY KEY (id));")
	ExecQuery(db, "CREATE TABLE IF NOT EXISTS public.authors (id uuid DEFAULT uuid_generate_v4 (), bio TEXT, author_name VARCHAR(255), papers uuid [], photo VARCHAR(255), psuedonym BOOL, created_at TIMESTAMPTZ, PRIMARY KEY (id));")
	ExecQuery(db, "CREATE TABLE IF NOT EXISTS public.files (id uuid DEFAULT uuid_generate_v4 (), cover_image VARCHAR(255), latest BOOL, pub_date TIMESTAMPTZ, source VARCHAR(255), url VARCHAR(255), version NUMERIC, created_at TIMESTAMPTZ, PRIMARY KEY (id));")
	ExecQuery(db, "CREATE TABLE IF NOT EXISTS public.papers (id uuid DEFAULT uuid_generate_v4 (), author uuid [], asset uuid [], description TEXT, excerpt TEXT, file uuid [], page_num SMALLINT, pretty_id SERIAL, title VARCHAR(255), created_at TIMESTAMPTZ, PRIMARY KEY (id));")
	ExecQuery(db, "CREATE TABLE IF NOT EXISTS public.users (id uuid DEFAULT uuid_generate_v4 (), name VARCHAR(255), email VARCHAR(255) NOT NULL UNIQUE, password TEXT NOT NULL, created_at TIMESTAMPTZ, PRIMARY KEY (id));")
	// Create connection tables
	ExecQuery(db, "CREATE TABLE IF NOT EXISTS public.author_paper (author_id uuid REFERENCES authors (id) ON UPDATE CASCADE, paper_id uuid REFERENCES papers (id) ON UPDATE CASCADE, CONSTRAINT author_paper_pkey PRIMARY KEY (author_id, paper_id));")

	// Insert default values
	ExecQuery(db, fmt.Sprintf("SET timezone = '%s';", conf.Base.Timezone))
	ExecQuery(db, fmt.Sprintf("INSERT INTO internal.config (config_name, config_value) VALUES ('schema_version', %s);", conf.Postgres.PGSchema))
}

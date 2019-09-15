package db

import (
	"database/sql"
	"fmt"
)

var configQuery = `
 	CREATE TABLE IF NOT EXISTS internal.config (
		id SERIAL,
		config_name VARCHAR(255),
		config_value VARCHAR(25),
		PRIMARY KEY (id)
	);
`

var assetsQuery = `
	CREATE TABLE IF NOT EXISTS public.assets (
		id uuid DEFAULT uuid_generate_v4 (),
		asset_name VARCHAR(255),
		logo VARCHAR(255),
		ticker VARCHAR(255),
		created_at TIMESTAMPTZ,
		PRIMARY KEY (id)
	);
`

var authorsQuery = `
	CREATE TABLE IF NOT EXISTS public.authors (
		id uuid DEFAULT uuid_generate_v4 (),
		author_name VARCHAR(255),
		bio TEXT,
		photo VARCHAR(255)
		psuedonym BOOL
		created_at TIMESTAMPTZ,
		PRIMARY KEY (id)
	);
`

var filesQuery = `
	CREATE TABLE IF NOT EXISTS public.files (
		id uuid DEFAULT uuid_generate_v4 (),
		cover_image VARCHAR(255),
		is_latest BOOL,
		page_num SMALLINT,
		pub_date TIMESTAMPTZ,
		source VARCHAR(255),
		url VARCHAR(255),
		version VARCHAR(255),
		created_at TIMESTAMPTZ,
		PRIMARY KEY (id)
	);
`

var papersQuery = `
	CREATE TABLE IF NOT EXISTS public.papers (
		id uuid DEFAULT uuid_generate_v4 (),
		description TEXT,
		excerpt TEXT,
		latest_version uuid,
		pretty_id SERIAL,
		title_primary VARCHAR(255),
		title_secondary VARCHAR(255),
		created_at TIMESTAMPTZ,
		PRIMARY KEY (id)
	);
`

var usersQuery = `
	CREATE TABLE IF NOT EXISTS public.users (
		id uuid DEFAULT uuid_generate_v4 (),
		email VARCHAR(255) NOT NULL UNIQUE,
		password TEXT NOT NULL,
		user_name VARCHAR(255),
		created_at TIMESTAMPTZ,
		PRIMARY KEY (id)
	);
`

var authorsPapersQuery = `
	CREATE TABLE IF NOT EXISTS public.author_paper (
		author_id uuid REFERENCES authors (id) ON UPDATE CASCADE,
		paper_id uuid REFERENCES papers (id) ON UPDATE CASCADE,
		CONSTRAINT author_paper_pkey
		PRIMARY KEY (author_id, paper_id)
	);
`

func schemaOne(db *sql.DB) {
	// Add extension to generate uuids
	ExecQuery(db, "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	// Create internal schema
	ExecQuery(db, "CREATE SCHEMA IF NOT EXISTS internal;")

	// Create required tables
	ExecQuery(db, configQuery)
	ExecQuery(db, assetsQuery)
	ExecQuery(db, authorsQuery)
	ExecQuery(db, filesQuery)
	ExecQuery(db, papersQuery)
	ExecQuery(db, usersQuery)
	// Create connection tables
	ExecQuery(db, authorsPapersQuery)

	// Insert default values
	ExecQuery(db, fmt.Sprintf("SET timezone = '%s';", conf.Base.Timezone))
	ExecQuery(db, fmt.Sprintf("INSERT INTO internal.config (config_name, config_value) VALUES ('schema_version', %s);", conf.Postgres.PGSchema))
}

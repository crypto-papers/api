package config

import "os"

// BaseConfig stores general configuration values
type BaseConfig struct {
	Timezone string
}

// GQLConfig stores configurations required by GraphQL
type GQLConfig struct {
	Port string
}

// PostgresConfig stores values needed to connect to postgres db
type PostgresConfig struct {
	PGUser     string
	PGPassword string
	PGName     string
	PGSchema   string
	PGSSL      string
}

// Config stores all API environmental variables
type Config struct {
	Base     BaseConfig
	GQL      GQLConfig
	Postgres PostgresConfig
}

// New generates a new Config struct containing environmental variable
func New() *Config {
	return &Config{
		Base: BaseConfig{
			Timezone: getEnv("TIMEZONE", "America/New_York"),
		},
		GQL: GQLConfig{
			Port: getEnv("GQL_PORT", "4000"),
		},
		Postgres: PostgresConfig{
			PGUser:     getEnv("PG_USER", ""),
			PGPassword: getEnv("PG_PASSWORD", ""),
			PGName:     getEnv("PG_DB_NAME", ""),
			PGSchema:   getEnv("PG_SCHEMA_VERSION", "1.0"),
			PGSSL:      getEnv("PG_SSL_MODE", ""),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

package config

import "os"

type GQLConfig struct {
	Port string
}

type PostgresConfig struct {
	PGUser     string
	PGPassword string
	PGName     string
	PGSSL      string
}

type Config struct {
	GQL      GQLConfig
	Postgres PostgresConfig
}

// New generates a new Config struct containing environmental variable
func New() *Config {
	return &Config{
		GQL: GQLConfig{
			Port: getEnv("GQL_PORT", "4000"),
		},
		Postgres: PostgresConfig{
			PGUser:     getEnv("PG_USER", ""),
			PGPassword: getEnv("PG_PASSWORD", ""),
			PGName:     getEnv("PG_DB_NAME", ""),
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

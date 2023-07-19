package database

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (cfg PostgresConfig) string() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode)
}

// func DefaultPostgresConfig() PostgresConfig {
// 	return PostgresConfig{
// 		Host:     os.Getenv("PSQL_HOST"),
// 		Port:     os.Getenv("PSQL_PORT"),
// 		User:     os.Getenv("PSQL_USER"),
// 		Password: os.Getenv("PSQL_PASSWORD"),
// 		Database: os.Getenv("PSQL_DATABASE"),
// 		SSLMode:  os.Getenv("PSQL_SSL_MODE"),
// 	}
// }
func NewPostgresConfig(host, port, user, password, database, sslmode string) PostgresConfig {
	return PostgresConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Database: database,
		SSLMode:  sslmode,
	}
}

func Open(config PostgresConfig) (*sql.DB, error) {
	db, err := sql.Open("pgx", config.string())
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

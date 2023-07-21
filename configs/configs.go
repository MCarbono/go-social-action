package configs

import (
	"go-social-action/infra/database"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	PSQL       database.PostgresConfig
	ServerPort string
}

func LoadEnvConfig(path string) (config, error) {
	var cfg config
	err := godotenv.Load(path)
	if err != nil {
		return cfg, err
	}
	cfg.PSQL = database.NewPostgresConfig(
		os.Getenv("PSQL_HOST"),
		os.Getenv("PSQL_PORT"),
		os.Getenv("PSQL_USER"),
		os.Getenv("PSQL_PASSWORD"),
		os.Getenv("PSQL_DATABASE"),
		os.Getenv("PSQL_SSL_MODE"),
	)
	cfg.ServerPort = os.Getenv("SERVER_PORT")
	return cfg, nil
}

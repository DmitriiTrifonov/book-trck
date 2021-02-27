package config

type Config struct {
	PostgresHost     string `envconfig:"POSTGRES_HOST" default:"localhost"`
	PostgresUser     string `envconfig:"POSTGRES_USER" default:"book-trck"`
	PostgresPassword string `envconfig:"POSTGRES_PASSWORD" default:"book-trck"`
	PostgresPort     string `envconfig:"POSTGRES_PORT" default:"5432"`
}

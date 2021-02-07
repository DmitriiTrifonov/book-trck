package config

type Config struct {
	PostgresHost    string `envconfig:"POSTGRES_HOST" default:"localhost"`
	PostgreUser     string `envconfig:"POSTGRES_USER" default:"book-trck"`
	PostgrePassword string `envconfig:"POSTGRES_PASSWORD" default:"book-trck"`
}

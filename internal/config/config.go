package config

import "fmt"

type Config struct {
	PostgresHost     string `envconfig:"POSTGRES_HOST" default:"localhost"`
	PostgresUser     string `envconfig:"POSTGRES_USER" default:"book-trck"`
	PostgresPassword string `envconfig:"POSTGRES_PASSWORD" default:"book-trck"`
	PostgresPort     string `envconfig:"POSTGRES_PORT" default:"5432"`
	PostgresDatabase string `envconfig:"POSTGRES_DATABASE" default:"book_trck"`
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable",
		c.PostgresHost, c.PostgresUser, c.PostgresPassword, c.PostgresPort, c.PostgresDatabase)
}

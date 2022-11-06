package configs

import (
	"fmt"

	"github.com/spf13/cast"
)

// PostgresConfig object
type PostgresConfig struct {
	Host     string `env:"DB_HOST"`
	Port     int    `env:"DB_PORT"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Name     string `env:"DB_NAME"`
}

// Dialect returns "postgres"
func (c PostgresConfig) Dialect() string {
	return "postgres"
}

// GetPostgresConnectionInfo returns Postgres URL string
func (c PostgresConfig) GetPostgresConnectionInfo() string {
	if c.Password == "" {
		return fmt.Sprintf(
			"host=%s port=%d user=%s dbname=%s sslmode=disable",
			c.Host, c.Port, c.User, c.Name)
	}
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.Name)
}

// GetPostgresConfig returns PostgresConfig object
func GetPostgresConfig() PostgresConfig {

	return PostgresConfig{
		Host:     GetEnvOrDefaultValue("DB_HOST", "0.0.0.0"),
		Port:     cast.ToInt(GetEnvOrDefaultValue("DB_PORT", "6543")),
		User:     GetEnvOrDefaultValue("DB_USER", "dev_user"),
		Password: GetEnvOrDefaultValue("DB_PASSWORD", "123test"),
		Name:     GetEnvOrDefaultValue("DB_NAME", "base_dev"),
	}
}

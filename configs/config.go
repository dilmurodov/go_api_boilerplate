package configs

import (
	"os"
)

const (
	prod = "production"
)

// Config object
type Config struct {
	Env       string         `env:"ENV"`
	Pepper    string         `env:"PEPPER"`
	HMACKey   string         `env:"HMAC_KEY"`
	Postgres  PostgresConfig `json:"postgres"`
	Mailgun   MailgunConfig  `json:"mailgun"`
	JWTSecret string         `env:"JWT_SIGN_KEY"`
	Host      string         `env:"APP_HOST"`
	Port      string         `env:"APP_PORT"`
	FromEmail string         `env:"EMAIL_FROM"`
}

// IsProd Checks if env is production
func (c Config) IsProd() bool {
	return c.Env == prod
}

// GetConfig gets all config for the application
func GetConfig() Config {
	return Config{
		Pepper:    GetEnvOrDefaultValue("PEPPER", ""),
		HMACKey:   GetEnvOrDefaultValue("HMAC_KEY", ""),
		JWTSecret: GetEnvOrDefaultValue("JWT_SIGN_KEY", ""),
		Postgres:  GetPostgresConfig(),
		Mailgun:   GetMailgunConfig(),
		Host:      GetEnvOrDefaultValue("APP_HOST", "0.0.0.0"),
		Port:      GetEnvOrDefaultValue("APP_PORT", "9000"),
		FromEmail: GetEnvOrDefaultValue("EMAIL_FROM", ""),
	}
}

func GetEnvOrDefaultValue(key string, def string) string {
	env, b := os.LookupEnv(key)
	if !b {
		return def
	}
	return env
}

//------------- Viper Implementation -------------------

// import (
// 	"log"
// 	"os"

// 	"github.com/spf13/viper"
// )

// func SetConfig() {
// 	if os.Getenv("ENV") != "PROD" {
// 		viper.SetConfigFile(".env")

// 		if err := viper.ReadInConfig(); err != nil {
// 			log.Fatalf("Error while reading config file %s", err)
// 		}
// 	} else {
// 		viper.AutomaticEnv()
// 	}
// }

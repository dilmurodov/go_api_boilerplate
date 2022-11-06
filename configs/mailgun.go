package configs

// MailgunConfig object
type MailgunConfig struct {
	APIKey string `env:"MAILGUN_API_KEY"`
	// PublicAPIKey string `env:"MAILGUN_PUBLIC_KEY"`
	Domain string `env:"MAILGUN_DOMAIN"`
}

// GetMailgunConfig get Mainlgun config object
func GetMailgunConfig() MailgunConfig {
	return MailgunConfig{
		// PublicAPIKey: os.Getenv("MAILGUN_PUBLIC_KEY"),
		APIKey: GetEnvOrDefaultValue("MAILGUN_API_KEY", ""),
		Domain: GetEnvOrDefaultValue("MAILGUN_DOMAIN", ""),
	}
}

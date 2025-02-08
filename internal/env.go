package internal

import (
	"flag"

	"github.com/abyanmajid/matcha/env"
	"github.com/abyanmajid/matcha/logger"
)

type EnvConfig struct {
	Origin           string `name:"ORIGIN" required:"true"`
	FrontendUrl      string `name:"FRONTEND_URL" required:"true"`
	DatabaseUrl      string `name:"DATABASE_URL" required:"true"`
	SmtpHost         string `name:"SMTP_HOST" required:"true"`
	SmtpPort         string `name:"SMTP_PORT" required:"true"`
	SmtpUser         string `name:"SMTP_USER" required:"true"`
	SmtpPassword     string `name:"SMTP_PASSWORD" required:"true"`
	EmailFrom        string `name:"EMAIL_FROM" required:"true"`
	JwtSecret        string `name:"JWT_SECRET" required:"true"`
	EncryptionSecret string `name:"ENCRYPTION_SECRET" required:"true"`
	EncryptionIv     string `name:"ENCRYPTION_IV" required:"true"`
}

func ConfigureEnv() *EnvConfig {
	dev := flag.Bool("dev", false, "Set the debug level (e.g., true or false)")

	flag.Parse()

	if *dev {
		if err := env.Dotenv(".env"); err != nil {
			logger.Fatal("Warning: No .env file found.")
		}
	}

	var config EnvConfig
	if err := env.Load(&config); err != nil {
		logger.Fatal("Error loading configuration: %s", err)
	}

	return &config
}
